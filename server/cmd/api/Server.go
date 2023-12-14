package server

import (
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hibiken/asynq"
	supa "github.com/nedpals/supabase-go"

	"github.com/hktrib/RetailGo/internal/ent"
	kv "github.com/hktrib/RetailGo/internal/redis"
	worker "github.com/hktrib/RetailGo/internal/tasks"
	"github.com/hktrib/RetailGo/internal/util"
	"github.com/hktrib/RetailGo/internal/weaviate"
	"github.com/hktrib/RetailGo/internal/webhook"
)

type Param string

// All Server operations are conducted using the following struct.

/*
	@Fields
		Router -> Chi's Multiplexers
		ClerkClient -> Clerk's Client (Authentication)
		DBClient -> Ent's Database Client
		TaskQueueClient -> Asynq Client (Task queue for event-triggered jobs)
		WeaviateClient -> A wrapper around Weaviate's Vector Database Client
		Cache -> A wrapper for Redis Key-Value Store / Cache
		TaskProducer -> An implementation of task producers for the task queue.
		Config -> Config related stuff
		Supabase -> A client for Supabase (used for image upload)
*/

type Server struct {
	Router          *chi.Mux
	ClerkClient     clerk.Client
	DBClient        *ent.Client
	TaskQueueClient *asynq.Client
	WeaviateClient  *weaviate.Weaviate
	Cache           *kv.Cache
	TaskProducer    worker.TaskProducer
	Config          *util.Config
	Supabase        *supa.Client
}

// Creates a new server.

func NewServer(
	clerkClient clerk.Client,
	entClient *ent.Client,
	taskQueueClient *asynq.Client,
	weaviateClient *weaviate.Weaviate,
	cache *kv.Cache,
	taskProducer worker.TaskProducer,
	config *util.Config,
) *Server {

	srv := &Server{}
	srv.Router = chi.NewRouter()
	srv.ClerkClient = clerkClient
	srv.DBClient = entClient
	srv.TaskQueueClient = taskQueueClient
	srv.WeaviateClient = weaviateClient
	srv.WeaviateClient.Start()
	srv.Cache = cache
	srv.TaskProducer = taskProducer
	srv.Config = config
	return srv
}

// Mount all mappings from route to handler in the mux.

func (s *Server) MountHandlers() {
	// Log all requests
	s.Router.Use(middleware.Logger)

	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           299, // Maximum value not ignored by any of major browsers
	}))

	// Route for Stripe webhook to handle transaction outcomes.
	s.Router.Route("/webhook/stripe", func(r chi.Router) {
		r.Post("/", func(writer http.ResponseWriter, request *http.Request) {
			webhook.StripeWebhookRouter(writer, request, s.Config.STRIPE_WEBHOOK_SECRET_CONNECTED, s.DBClient)
		})
		r.Post("/connect", func(writer http.ResponseWriter, request *http.Request) {
			webhook.StripeWebhookRouter(writer, request, s.Config.STRIPE_WEBHOOK_SECRET_CONNECTED, s.DBClient)
		})

	})

	// POST | Stages Clerk User Data via loopback to /IsOwnerCreateHandle
	s.Router.Post("/clerkwebhook", webhook.HandleClerkWebhook)

	s.Router.Get("/", func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(([]byte)("OK"))
	})

	s.Router.Group(func(r chi.Router) {

		// Add necessary middleware for protecting owner and store create routes
		r.Route("/create", func(r chi.Router) {

			// MIDDLEWARE | Validates Clerk signature before Create Store/Owner
			// r.Use(s.ProtectStoreAndOwnerCreation)

			// Handles User Data for Potential Store Creation
			r.Group(func(r chi.Router) {

				// MIDDLEWARE | Stages USER data for store creation in Redis
				r.Use(s.IsOwnerCreateHandle)

				// POST | Creates a new user
				r.Post("/user", s.UserCreate)

			})

			// Handles Store Creation Requests
			r.Group(func(r chi.Router) {

				// MIDDLEWARE | Fetches Staged User Data from Redis
				r.Use(s.StoreCreateHandle)

				// POST | Creates a new Store -> using Staged Owner Data
				r.Post("/store", s.CreateStore)

			})
		})
	})

	s.Router.Route("/store", func(r chi.Router) {
		r.Route("/{store_id}", func(r chi.Router) {

			// store access validation middleware
			//	r.Use(s.ValidateStoreAccess)

			// Onboarding Route(s)
			r.Group(func(r chi.Router) {
				// owner validation middleware
				//	r.Use(s.ValidateOwner)

				// Get a store by ID
				r.Get("/onboarding", s.HandleOnboarding)

			})

			// Inventory Route(s)
			r.Route("/inventory", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					// MIDDLEWARE | Owner validation
					//		r.Use(s.ValidateOwner)

					// POST | Create Item in Inventory
					r.Post("/create", s.InvCreate)

					// DELETE | Delete Item in Inventory
					r.Delete("/", s.InvDelete)
				})

				// GET | Gets all items in Inventory
				r.Get("/", s.InvRead)

				// POST | Updates Item in Inventory
				r.Post("/update", s.InvUpdate)

				// POST | Updates Item's Photo URL in Inventory
				r.Post("/updatephoto", s.InvUpdatePhoto)

			})

			// Category Route(s)
			r.Route("/category", func(r chi.Router) {

				// GET | Get's all Items in the category
				r.Get("/", s.CatItemList)

				// Category CRUD Route(s)
				r.Group(func(r chi.Router) {

					// MIDDLEWARE | Owner validation
					r.Use(s.ValidateOwner)

					// POST | creating new category
					r.Post("/", s.CatCreate)

					// DELETE | deleting existing category
					r.Delete("/", s.CatDelete)

				})

				// Category Item(s) CRUD Route(s)
				r.Route("/{category_id}", func(r chi.Router) {
					r.Group(func(r chi.Router) {

						// MIDDLEWARE | Owner validation
						//	r.Use(s.ValidateOwner)

						// GET | All items from category
						r.Get("/", s.CatItemRead)

						// POST | Add an item to a category
						r.Post("/", s.CatItemAdd)

					})
				})

				// POST | Update Item in Category
				r.Post("/update", s.InvUpdate)
			})

			// Staff Routes
			r.Route("/staff", func(r chi.Router) {

				// Get | All staff/employees from store
				r.Get("/", s.GetAllEmployees)

				// Post | Send invitation to prospective employee
				r.Post("/invite", s.SendInviteEmail)
			})

			// POS Routes
			r.Route("/pos", func(r chi.Router) {
				r.Post("/checkout", s.StoreCheckout)
				r.Get("/info", s.GetPosInfo)
			})

			// Users/
			r.Route("/users", func(r chi.Router) {
				r.Get("/", s.GetStoreUsers)
			})
		})

		// Get | Store by UUID
		r.Get("/uuid/{uuid}", s.GetStoreByUUID)

		// Get | Store by User's Clerk ID
		r.Get("/", s.GetStoresByClerkId)
	})

	s.Router.Route("/user", func(r chi.Router) {

		// Checks if a user has a store
		r.Get("/store", s.UserHasStore)

		r.Route("/{user_id}", func(r chi.Router) {
			// Delete a user by ID
			r.Delete("/", s.UserDelete)

			// Update a user by ID
			r.Put("/", s.userUpdate)

			// Get a user by ID
			r.Get("/", s.UserQuery)
		})

		// Join a store for a user
		r.Post("/join", s.UserJoinStore)

		// Create a new user
		r.Post("/", s.UserCreate)

		// Additional user creation endpoint
		r.Post("/add", s.userAdd)
	})

}
