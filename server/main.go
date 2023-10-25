package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

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

	client := Open(config.DBSource)
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
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
