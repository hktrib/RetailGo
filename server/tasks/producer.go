package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

/*
 */
type TaskProducer interface {
	TaskOwnerCreationCheck(
		ctx context.Context,
		ownerEmailID *string,
		opts ...asynq.Option,
	) error
}

type RedisProducer struct {
	messageQueueClient *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskProducer {
	client := asynq.NewClient(redisOpt)
	return &RedisProducer{
		messageQueueClient: client,
	}
}
