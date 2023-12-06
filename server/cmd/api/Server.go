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

/*
	@Fields
		Router -> Chi's Multiplexers
		Client -> Ent's Database Client
		Config -> Config related stuff

*/

type Param string

type Server struct {
	Router          *chi.Mux
	ClerkClient     clerk.Client
	DBClient        *ent.Client
	TaskQueueClient *asynq.Client
	WeaviateClient  *weaviate.Weaviate
	Cache           *kv.Cache
	TaskProducer    worker.TaskProducer
	Config          *util.Config
	Supabase 		*supa.Client
}

// Define a type for Item-Change Requests (Create, Update, Delete)

func NewServer(
	clerkClient clerk.Client,
	entClient *ent.Client,
	taskQueueClient *asynq.Client,
	weaviateClient *weaviate.Weaviate,
	cache *kv.Cache,
	taskProducer worker.TaskProducer,
	config *util.Config,
	supabase *supa.Client,
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
	srv.Supabase = supabase
	
	return srv
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           299, // Maximum value not ignored by any of major browsers
	}))

	s.Router.Post("/webhook", func(writer http.ResponseWriter, request *http.Request) {
		s.StripeWebhookRouter(writer, request)
	})
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
			r.Route("/inventory", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					//		r.Use(s.ValidateOwner)         // add owner validation validation
					r.Post("/create", s.InvCreate) //
					r.Delete("/", s.InvDelete)     //
				})
				r.Get("/", s.InvRead)          //
				r.Post("/update", s.InvUpdate) //
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
				r.Get("/checkout", s.StoreCheckout)
				r.Get("/info", s.GetPosInfo)
			})
			r.Route("/users", func(r chi.Router) {
				r.Get("/", s.GetStoreUsers)
			})
		})
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
