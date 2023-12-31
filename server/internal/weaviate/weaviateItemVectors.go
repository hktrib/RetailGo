package weaviate

import (
	"context"
	"fmt"
	"time"

	worker "github.com/hktrib/RetailGo/internal/tasks"
)

func (weaviate *Weaviate) DoVectorize(ctx context.Context, taskProducer worker.TaskProducer) error {

	for {
		err := taskProducer.ProduceTaskUpdateVectors(ctx, time.Minute)
		if err != nil {
			fmt.Println("Error producing update vectors task", err)
			return err
		}
		time.Sleep(10 * time.Minute)
	}
}
