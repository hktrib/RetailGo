package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"

	server "github.com/hktrib/RetailGo/cmd/api"
	"github.com/hktrib/RetailGo/internal/ent"
	kvRedis "github.com/hktrib/RetailGo/internal/redis"
	worker "github.com/hktrib/RetailGo/internal/tasks"
	"github.com/hktrib/RetailGo/internal/util"
	weaviate "github.com/hktrib/RetailGo/internal/weaviate"
	"github.com/hktrib/RetailGo/internal/webhook"
)

// var log = util.NewLogger()

func runTaskConsumer(redisOptions *asynq.RedisClientOpt, dbClient *ent.Client, clerkclient clerk.Client, config *util.Config) {
	taskConsumer := worker.NewRedisTaskConsumer(*redisOptions, dbClient, clerkclient, config)
	err := taskConsumer.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start task consumer")
	}
}

func main() {
	fmt.Println("Starting server...")
	config, err := util.LoadConfig()
	if err != nil {
		panic(err)
	}

	stripe.Key = config.STRIPE_SK

	taskQueueOptions := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%s", config.REDIS_HOSTNAME, config.REDIS_PORT),
		Password: config.REDIS_PASSWORD,
		DB:       1,
	}

	cacheOptions := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.REDIS_HOSTNAME, config.REDIS_PORT),
		Password: config.REDIS_PASSWORD,
		DB:       0,
	}

	clerkClient, err := clerk.NewClient(config.CLERK_SK)
	if err != nil {
		panic(err)
	}

	taskProducer := worker.NewRedisTaskProducer(taskQueueOptions)
	fmt.Printf("checkpoint 1")
	cache := kvRedis.NewCache(context.Background(), cacheOptions, 1*time.Minute)
	defer cache.Client.Close()

	taskQueueClient := asynq.NewClient(taskQueueOptions)
	defer taskQueueClient.Close()

	entClient := util.Open(&config)
	defer entClient.Close()
	fmt.Printf("checkpoint 2")
	weaviateClient := weaviate.NewWeaviate(context.Background())

	if err != nil {
		panic(err)
	}

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}
	println("checkpoint 3")

	go weaviateClient.DoVectorize(context.Background(), taskProducer)

	go func() {
		injectActiveSession := clerk.WithSessionV2(clerkClient)

		srv := server.NewServer(clerkClient, entClient, taskQueueClient, weaviateClient, cache, taskProducer, &config)
		srv.Router.Use(injectActiveSession)

		srv.MountHandlers()

		webhook.Config = &config
		webhook.ClerkClient = clerkClient
		println("checkpoint 4")

		go runTaskConsumer(&taskQueueOptions, entClient, clerkClient, &config)
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", config.SERVER_ADDRESS), srv.Router)
		println("checkpoint 5")

		log.Debug().Msg("Deploy Msg: starting server!")
		if err != nil {
			log.Fatal().Err(err).Msg("failed in starting server")
		}
	}()

	// Makes sure we wait for the go routine running
	select {}
}
