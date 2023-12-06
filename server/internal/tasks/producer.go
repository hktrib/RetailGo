package worker

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
)

/*
 */
type TaskProducer interface {
	ProduceTaskOwnerCreationCheck(
		ctx context.Context,
		ownerEmailID *string,
		processInTime time.Duration,
		opts ...asynq.Option,
	) error
	ProduceTaskUpdateVectors(
		ctx context.Context,
		processInTime time.Duration,
		opts ...asynq.Option,
	) error
}

type RedisProducer struct {
	messageQueueClient *asynq.Client
}

func NewRedisTaskProducer(redisOpt asynq.RedisClientOpt) TaskProducer {
	client := asynq.NewClient(redisOpt)
	return &RedisProducer{
		messageQueueClient: client,
	}
}
