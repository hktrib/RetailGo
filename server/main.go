package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect"
	"github.com/hktrib/RetailGo/ent"
	server "github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func Open(config *util.Config) *ent.Client {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	client := Open(&config)
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := server.NewServer(client, &config)
	server.MountHandlers()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), server.Router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Makes sure we wait for the go routine running
	select {}
}
