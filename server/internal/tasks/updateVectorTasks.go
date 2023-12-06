package worker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hibiken/asynq"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/transactions"
	"github.com/rs/zerolog/log"
)

// A list of applicable task types.
const (
	TaskUpdateVectors = "vectors:update"
)

// Need to marshall items to JSON, send them to the Python Server, receive the response, read the body, read the ids that were not vectorized, and set them
func vectorize(unvectorizedItems []*ent.Item) ([]int, error) {
	REC_SERVER_URL := "https://recommendation-server-production.up.railway.app"
	// Marshall items to JSON
	requestBody, err := json.Marshal(unvectorizedItems)
	var ids []int

	if err != nil {
		return ([]int{}), err
	}

	response, err := http.Post(REC_SERVER_URL+"vectorizeItems", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return ([]int{}), err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return ([]int{}), err
	}

	err = json.Unmarshal(responseBody, &ids)

	if err != nil {
		return ([]int{}), err
	}

	return ids, nil
}

func (rp *RedisProducer) ProduceTaskUpdateVectors(ctx context.Context, processInTime time.Duration, opts ...asynq.Option) error {
	// Create a new asynq task with the correct type, and passing down the opts
	task := asynq.NewTask(TaskUpdateVectors, nil, opts...)
	// Enqueue the task on the messageQueue, with a processInTime
	info, err := rp.messageQueueClient.Enqueue(task, asynq.ProcessIn(processInTime))
	if err != nil {
		return fmt.Errorf("failed to enqueue update vector task: %w", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")

	return nil
}

func (rc *RedisConsumer) ConsumeTaskUpdateVectors(ctx context.Context, task *asynq.Task) error {
	// Run the transaction to query and set to vectorized all unvectorized items
	unvectorizedItems, err := transactions.QueryUnvectorizedItemsAndMarkVectorized(ctx, rc.dbClient)

	if err != nil {
		return fmt.Errorf("failed to query/update unvectorized items to vectorized: %w", err)
	}

	// Send all unvectorized items to Python Server
	// Add to ENVVARs + Config.go, but for now, use the actual URL.
	// Need to marshall items to JSON, send them to the Python Server, receive the response, read the body, read the ids that were not vectorized, and set them
	itemsFailedToVectorize, err := vectorize(unvectorizedItems)
	if err != nil {
		return fmt.Errorf("failed to vectorize: %w", err)
	}

	// Based on result, update the database.
	for itemID := range itemsFailedToVectorize {
		failedItem, err := rc.dbClient.Item.Query().Where(item.ID(itemID)).Only(ctx)
		if err != nil {
			continue
		}

		_, _ = failedItem.Update().SetVectorized(false).SetNumberSoldSinceUpdate(0).Save(ctx)
	}

	return nil
}
