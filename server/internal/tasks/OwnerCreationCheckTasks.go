package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/user"
)

// A list of applicable task types.
const (
	TaskOwnerCreationCheck = "store|owner:create"
)

func (rp *RedisProducer) ProduceTaskOwnerCreationCheck(ctx context.Context, ownerEmailID *string, processInTime time.Duration, opts ...asynq.Option) error {
	payload, err := json.Marshal(ownerEmailID)
	if err != nil {
		return err
	}
	task := asynq.NewTask(TaskOwnerCreationCheck, payload, opts...)

	info, err := rp.messageQueueClient.Enqueue(task, asynq.ProcessIn(processInTime))
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (rc *RedisConsumer) ConsumeTaskOwnerCreationCheck(ctx context.Context, task *asynq.Task) error {
	var ownerEmail string
	if err := json.Unmarshal(task.Payload(), &ownerEmail); err != nil {
		return fmt.Errorf("error: unable to marshal : %w", asynq.SkipRetry)
	}

	_, err := rc.dbClient.User.Query().Where(user.Email(ownerEmail)).Only(ctx)
	if err != nil {
		if _, isMyErrorType := err.(*ent.NotFoundError); isMyErrorType {

			log.Info().Msg("Sending reminder Email to user")

			return fmt.Errorf("error: User has no store Association yet!!!%w", err)
		} else {

			// Somehow we have 2 records of the user with the same email!!!
			return fmt.Errorf("error: 2 Records of same user!!! %w", err)
		}
	}

	return nil
}
