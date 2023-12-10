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

type ItemBatch struct {
	Items []ent.Item
}

func getSales(unvectorizedItems []ent.Item) (map[int]int, []int) {
	unvectorizedSales := make(map[int]int)
	var unvectorizedIds []int

	for _, item := range unvectorizedItems {
		// unvectorizedSales = append(unvectorizedSales, item.NumberSoldSinceUpdate)
		unvectorizedSales[item.ID] = item.NumberSoldSinceUpdate
		unvectorizedIds = append(unvectorizedIds, item.ID)
	}

	return unvectorizedSales, unvectorizedIds
}

// Need to marshall items to JSON, send them to the Python Server, receive the response, read the body, read the ids that were not vectorized, and set them
func vectorize(unvectorizedItems []ent.Item) ([]int, error) {
	REC_SERVER_URL := "https://recommendation-server-production.up.railway.app"
	// Marshall items to JSON
	itemBatch := new(ItemBatch)
	itemBatch.Items = unvectorizedItems

	requestBody, err := json.Marshal(itemBatch)
	fmt.Println("Request Body:", (string)(requestBody))

	var ids []int

	if err != nil {
		return ([]int{}), err
	}

	response, err := http.Post(REC_SERVER_URL+"/vectorizeItems", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return ([]int{}), err
	}

	if response.StatusCode == 500 {
		fmt.Println("Failed to vectorize: Internal Server Error.")
		return ([]int{}), err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	fmt.Printf("ResponseBody: %s\n", responseBody)

	if err != nil {
		return ([]int{}), err
	}

	err = json.Unmarshal(responseBody, &ids)

	if err != nil {
		return ([]int{}), err
	}

	return ids, nil
}

func rollbackSales(ctx context.Context, dbClient *ent.Client, itemsFailedToVectorize []int, idToNumberSoldSinceUpdate map[int]int) {
	for itemID := range itemsFailedToVectorize {
		failedItem, err := dbClient.Item.Query().Where(item.ID(itemID)).Only(ctx)
		if err != nil {
			continue
		}

		_, _ = failedItem.Update().SetVectorized(false).SetNumberSoldSinceUpdate(failedItem.NumberSoldSinceUpdate + idToNumberSoldSinceUpdate[itemID]).Save(ctx)
	}
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
	idToNumberSoldSinceUpdate, idsUnvectorized := getSales(unvectorizedItems)

	if err != nil {
		rollbackSales(ctx, rc.dbClient, idsUnvectorized, idToNumberSoldSinceUpdate)
		return fmt.Errorf("failed to query/update unvectorized items to vectorized: %w", err)
	} else if len(unvectorizedItems) == 0 {
		fmt.Println("no items to vectorize.")
		return nil
	}

	// Send all unvectorized items to Python Server
	itemsFailedToVectorize, err := vectorize(unvectorizedItems)
	if err != nil {
		rollbackSales(ctx, rc.dbClient, idsUnvectorized, idToNumberSoldSinceUpdate)
		return fmt.Errorf("failed to vectorize: %w", err)
	}

	// Based on result, update the database.
	rollbackSales(ctx, rc.dbClient, itemsFailedToVectorize, idToNumberSoldSinceUpdate)

	return nil
}
