package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	server "github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	client := util.Open(&config)
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := server.NewServer(client, &config)
	srv.MountHandlers()

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), srv.Router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Makes sure we wait for the go routine running
	select {}
}
