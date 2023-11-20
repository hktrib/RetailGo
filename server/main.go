package main

import (
	"context"
	"fmt"
	"github.com/stripe/stripe-go/v76"
	"net/http"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/hibiken/asynq"
	server "github.com/hktrib/RetailGo/cmd/api"
	"github.com/hktrib/RetailGo/internal/ent"
	kvRedis "github.com/hktrib/RetailGo/internal/redis"
	worker "github.com/hktrib/RetailGo/internal/tasks"
	"github.com/hktrib/RetailGo/internal/util"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
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
	config, err := util.LoadConfig(".")
	fmt.Printf("Test")
	if err != nil {
		panic(err)
	}
	stripe.Key = "sk_test_51ODz7pHWQUATs9zV4fWYLtRag0GwwLPticrlOe5FqicEWwdnWUlsZkRh90o1YOkt3qsOduJQNSbbUJupkm4i9xLm00hcffWjDm"

	taskQueueOptions := asynq.RedisClientOpt{
		Addr: fmt.Sprintf("0.0.0.0:%v", config.RedisAddress),
		DB:   1,
	}

	cacheOptions := &redis.Options{
		Addr:     fmt.Sprintf("0.0.0.0:%v", config.RedisAddress),
		Password: "",
		DB:       0,
	}

	clerkClient, err := clerk.NewClient(config.ClerkSK)
	if err != nil {
		panic(err)
	}

	taskProducer := worker.NewRedisTaskProducer(taskQueueOptions)

	cache := kvRedis.NewCache(context.Background(), cacheOptions, 1*time.Minute)
	defer cache.Client.Close()

	taskQueueClient := asynq.NewClient(taskQueueOptions)
	defer taskQueueClient.Close()

	entClient := util.Open(&config)
	defer entClient.Close()

	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed creating schema resources")
	}

	go func() {
		injectActiveSession := clerk.WithSessionV2(clerkClient)

		srv := server.NewServer(clerkClient, entClient, taskQueueClient, cache, taskProducer, &config)
		srv.Router.Use(injectActiveSession)

		srv.MountHandlers()
		err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), srv.Router)
		if err != nil {
			log.Fatal().Err(err).Msg("failed in starting server")
		}

		go runTaskConsumer(&taskQueueOptions, entClient, clerkClient, &config)

	}()

	// Makes sure we wait for the go routine running
	select {}
}
