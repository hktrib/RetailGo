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
			sk := s.Config.STRIPE_WEBHOOK_SECRET
			webhook.StripeWebhookRouter(writer, request, sk, s.DBClient)
		})
		r.Post("/connect", func(writer http.ResponseWriter, request *http.Request) {
			sk := s.Config.STRIPE_WEBHOOK_SECRET_CONNECTED

			webhook.StripeWebhookRouter(writer, request, sk, s.DBClient)
		})

	})

	// Route for Clerk Webhook.
	s.Router.Post("/clerkwebhook", webhook.HandleClerkWebhook)
	s.Router.Get("/", func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(([]byte)("OK"))
	})

	s.Router.Group(func(r chi.Router) {

		// Add necessary middleware for protecting owner and store create routes
		r.Route("/create", func(r chi.Router) {
			// r.Use(s.ProtectStoreAndOwnerCreation)
			r.Group(func(r chi.Router) {
				r.Use(s.IsOwnerCreateHandle)
				r.Post("/user", s.UserCreate)
			})
			r.Group(func(r chi.Router) {
				r.Use(s.StoreCreateHandle)
				r.Post("/store", s.CreateStore)
			})
		})
	})

	s.Router.Route("/store", func(r chi.Router) {
		r.Route("/{store_id}", func(r chi.Router) {
			//	r.Use(s.ValidateStoreAccess) // add [Employee || Owner] validation
			r.Group(func(r chi.Router) {
				//	r.Use(s.ValidateOwner)                   // add owner validation validation
				r.Get("/onboarding", s.HandleOnboarding) // Get a store by ID)
			})
			r.Route("/inventory", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					//		r.Use(s.ValidateOwner)         // add owner validation validation
					r.Post("/create", s.InvCreate) //
					r.Delete("/", s.InvDelete)     //
				})
				r.Get("/", s.InvRead)                    //
				r.Post("/update", s.InvUpdate)           //
				r.Post("/updatephoto", s.InvUpdatePhoto) //
			})
			r.Route("/category", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(s.ValidateOwner)     // add owner validation validation
					r.Post("/", s.CatCreate)   //
					r.Delete("/", s.CatDelete) //
				})
				r.Route("/{category_id}", func(r chi.Router) {
					r.Group(func(r chi.Router) {
						//	r.Use(s.ValidateOwner)         // add owner validation validation
						//	r.Delete("/", s.CatItemDelete) //
						r.Get("/", s.CatItemRead) //
						r.Post("/", s.CatItemAdd) //
					})
				})

				r.Get("/", s.CatItemList)      //
				r.Post("/update", s.InvUpdate) //

			})

			r.Route("/staff", func(r chi.Router) {
				r.Get("/", s.GetAllEmployees) //
				r.Post("/invite", s.SendInviteEmail)
			})

			r.Route("/pos", func(r chi.Router) {
				r.Post("/checkout", s.StoreCheckout)
				r.Get("/info", s.GetPosInfo)
			})
			r.Route("/users", func(r chi.Router) {
				r.Get("/", s.GetStoreUsers)
			})
		})
		r.Get("/uuid/{uuid}", s.GetStoreByUUID) //

		r.Get("/", s.GetStoresByClerkId)
	})

	s.Router.Route("/user", func(r chi.Router) {
		r.Get("/store", s.UserHasStore) // Checks if a user has a store
		r.Route("/{user_id}", func(r chi.Router) {
			r.Delete("/", s.UserDelete) // Delete a user by ID
			r.Put("/", s.userUpdate)    // Update a user by ID
			r.Get("/", s.UserQuery)     // Get a user by ID
		})
		r.Post("/join", s.UserJoinStore) // Join a store for a user
		r.Post("/", s.UserCreate)        // Create a new user
		r.Post("/add", s.userAdd)        // Additional user creation endpoint
	})

}
