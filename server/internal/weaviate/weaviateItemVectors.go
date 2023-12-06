package weaviate

import (
	"context"
	"time"

	worker "github.com/hktrib/RetailGo/internal/tasks"
)

func (weaviate *Weaviate) DoVectorize(ctx context.Context, taskProducer worker.TaskProducer) error {

	for {
		err := taskProducer.ProduceTaskUpdateVectors(ctx, time.Hour)
		if err != nil {
			return err
		}
		time.Sleep(time.Hour)
	}
}