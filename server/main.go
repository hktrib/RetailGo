package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"

	_ "github.com/lib/pq"
)

<<<<<<< HEAD
func Open(config *util.Config) *ent.Client {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

=======
>>>>>>> c79e5e84257577c1ed2958b25dea340f6cf088ce
func setupRoutes(r chi.Router, routes *routes.RouteHandler) {
	// Public Routes
	r.Group(func(r chi.Router) {
		r.Get("/", routes.HelloWorld)
	})
	// Private Routes
	// r.Group(func(r chi.Router) {
	//     r.Use(AuthMiddleware)
	//     r.Post("/manage", CreateAsset)
	// })
	r.Route("/store", func(r chi.Router) {
		r.Route("/{store_id}", func(r chi.Router) {
			r.Use(routes.ValidateStore) // add user validation
			r.Route("/inventory", func(r chi.Router) {
				r.Get("/", routes.InvRead)          //
				r.Post("/update", routes.InvUpdate) //
				r.Delete("/", routes.InvDelete)     //
				r.Post("/create", routes.InvCreate) //
			})
		})
	})
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

<<<<<<< HEAD
	client := Open(&config)
=======
	client, err := ent.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
>>>>>>> c79e5e84257577c1ed2958b25dea340f6cf088ce
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	routeHandler := routes.RouteHandler{
		Client: client,
		Config: &config,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	setupRoutes(r, &routeHandler)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Makes sure we wait for the go routine running
	select {}
}
