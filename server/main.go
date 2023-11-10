package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-redis/redis"
	"github.com/hibiken/asynq"
	server "github.com/hktrib/RetailGo/routes"
	"github.com/hktrib/RetailGo/util"
	_ "github.com/redis/go-redis/v9"

	_ "github.com/lib/pq"
)

func main() {

	// ctx := context.Background()

	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	kvStoreClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%v", config.RedisAddress),
		Password: "",
		DB:       0,
	})
	defer kvStoreClient.Close()

	messageQueueClient := asynq.NewClient(asynq.RedisClientOpt{
		Addr: fmt.Sprintf("localhost:%v", config.RedisAddress),
		DB:   1,
	})
	defer messageQueueClient.Close()

	clerkClient, err := clerk.NewClient(config.ClerkSK)
	if err != nil {
		panic(err)
	}

	injectActiveSession := clerk.WithSessionV2(clerkClient)

	entClient := util.Open(&config)
	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	srv := server.NewServer(clerkClient, entClient, &config)
	srv.Router.Use(injectActiveSession)

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
