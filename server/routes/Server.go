package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/util"
)

/*
	@Fields
		Router -> Chi's Multiplexers
		Client -> Ent's Database Client
		Config -> Config related stuff

*/

type Server struct {
	Router *chi.Mux
	Client *ent.Client
	Config *util.Config
}

func NewServer(client *ent.Client, config *util.Config) *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	s.Client = client
	s.Config = config
	return s
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
			r.Use(s.ValidateStore) // add user validation
			r.Route("/inventory", func(r chi.Router) {
				r.Get("/", s.InvRead)          //
				r.Post("/update", s.InvUpdate) //
				r.Delete("/", s.InvDelete)     //
				r.Post("/create", s.InvCreate) //
			})
		})
	})

}
