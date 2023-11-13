package server

import (
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hibiken/asynq"
	"github.com/hktrib/RetailGo/ent"
	kv "github.com/hktrib/RetailGo/redis"
	worker "github.com/hktrib/RetailGo/tasks"
	"github.com/hktrib/RetailGo/util"
)

/*
	@Fields
		Router -> Chi's Multiplexers
		Client -> Ent's Database Client
		Config -> Config related stuff

*/

type Server struct {
	Router          *chi.Mux
	ClerkClient     clerk.Client
	DBClient        *ent.Client
	TaskQueueClient *asynq.Client
	Cache           *kv.Cache

	TaskProducer worker.TaskProducer
	Config       *util.Config
}

func NewServer(
	clerkClient clerk.Client,
	entClient *ent.Client,
	taskQueueClient *asynq.Client,
	cache *kv.Cache,
	taskProducer worker.TaskProducer,
	config *util.Config,
) *Server {

	srv := &Server{}
	srv.Router = chi.NewRouter()
	srv.ClerkClient = clerkClient
	srv.DBClient = entClient
	srv.TaskQueueClient = taskQueueClient
	srv.Cache = cache

	srv.TaskProducer = taskProducer
	srv.Config = config
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

	s.Router.Group(func(r chi.Router) {
		r.Get("/", s.HelloWorld)
	})

	s.Router.Route("/store", func(r chi.Router) {
		r.Route("/{store_id}", func(r chi.Router) {
			r.Use(s.ValidateStoreAccess) // add [Employee || Owner] validation
			r.Route("/inventory", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(s.ValidateOwner)         // add owner validation validation
					r.Post("/create", s.InvCreate) //
					r.Delete("/", s.InvDelete)     //
				})
				r.Get("/", s.InvRead)          //
				r.Post("/update", s.InvUpdate) //
			})
		})
	})

}
