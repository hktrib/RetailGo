package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	db "github.com/hktrib/RetailGo/db/sqlc"
	"github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"

	_ "github.com/lib/pq"
)

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
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	routes := routes.RouteHandler{
		Store:  &store,
		Config: &config,
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	setupRoutes(r, &routes)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Makes sure we wait for the go routine running
	select {}
}
