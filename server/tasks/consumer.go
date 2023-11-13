package worker

import (
	"context"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/hibiken/asynq"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/util"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

const (
	CriticalQueue = "critical"
	DefaultQueue  = "default"
)

type TaskConsumer interface {
	Start() error
	ConsumeTaskOwnerCreationCheck(ctx context.Context, task *asynq.Task) error
}

type RedisConsumer struct {
	server      *asynq.Server
	dbClient    *ent.Client
	clerkClient clerk.Client
	config      *util.Config
}

func NewRedisTaskConsumer(redisOpt asynq.RedisClientOpt, dbClient *ent.Client, clerkclient clerk.Client, config *util.Config) TaskConsumer {
	logger := util.NewLogger()
	redis.SetLogger(logger)

	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				CriticalQueue: 10,
				DefaultQueue:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Error().Err(err).Str("type", task.Type()).
					Bytes("payload", task.Payload()).Msg("processer task failed")
			}),
			Logger: logger,
		},
	)

	return &RedisConsumer{
		server:      server,
		dbClient:    dbClient,
		clerkClient: clerkclient,
		config:      config,
	}
}

func (p *RedisConsumer) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskOwnerCreationCheck, p.ConsumeTaskOwnerCreationCheck)
	return p.server.Start(mux)
}
