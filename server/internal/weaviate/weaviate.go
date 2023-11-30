package weaviate

import (
	"context"
	"log"

	server "github.com/hktrib/RetailGo/cmd/api"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/util"
	weaviateClient "github.com/weaviate/weaviate-go-client/v4/weaviate"
	weaviateAuth "github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate/entities/models"
)

type Weaviate struct {
	Client *weaviateClient.Client
}

// Define a type for Item-Change Requests (Create, Update, Delete)

type ItemChange struct {
	Item          ent.Item
	Mode          string
	UpdatedFields server.UpdatedFields
}

func (weaviate *Weaviate) Start() {
	config, err := util.LoadConfig()

	if err != nil {
		panic(err)
	}

	weaviateConfig := weaviateClient.Config{
		Host:       config.WEAVIATE_HOSTNAME,
		Scheme:     "https",
		AuthConfig: weaviateAuth.ApiKey{Value: config.WEAVIATE_SK},
		Headers:    nil,
	}

	weaviate.Client, err = weaviateClient.NewClient(weaviateConfig)

	if err != nil {
		panic(err)
	}

	classObj := &models.Class{
		Class:        "Item",
		Vectorizer:   "none",
		ModuleConfig: map[string]interface{}{},
	}

	err = weaviate.Client.Schema().ClassCreator().WithClass(classObj).Do(context.Background())

	if err != nil {
		panic(err)
	}
}

func (weaviate *Weaviate) DispatchChanges(ctx context.Context, itemChange *ItemChange) {
	// NOTE TO SELF: FULL UPDATE SHOULD BE SCHEDULED, NOT JUST THE VECTORIZATION IN ALL CASES. LET THE LAMBDA HANDLE ALL OF THAT.

	if itemChange.Mode == "Create" {
		// Update literal details on Weaviate.
	} else if itemChange.Mode == "Update" {
		// Update literal details on Weaviate.

		// If only unvectorized fields are changed, exit.
	} else if itemChange.Mode == "Delete" {
		// Use the client to delete the item from Weaviate and exit.
		err := weaviate.Client.
			Data().
			Deleter().
			WithClassName("Item").
			WithID(itemChange.Item.WeaviateID).
			Do(ctx)

		if err == nil {
			return
		}

	} else {
		log.Fatal("Invalid Item Change Mode.")
	}

	// Send the ItemChange to Redis.

}

func (weaviate *Weaviate) ItemChangeHandler(ctx context.Context, itemChangeChannel chan *ItemChange) {
	for itemChange := range itemChangeChannel {
		go weaviate.DispatchChanges(ctx, itemChange)
	}
	select {}
}
