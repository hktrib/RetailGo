package weaviate

import (
	"context"
	"log"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/util"
	weaviateClient "github.com/weaviate/weaviate-go-client/v4/weaviate"
	weaviateAuth "github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
)

type Weaviate struct {
	Client            *weaviateClient.Client
	ctx               context.Context
	itemChangeChannel chan ItemChange
}

func ChangesVectorizedProperties(itemChange ItemChange) bool {
	return (itemChange.UpdatedFields.Name || itemChange.UpdatedFields.CategoryName || itemChange.UpdatedFields.Price || itemChange.UpdatedFields.Photo || itemChange.UpdatedFields.DateLastSold || itemChange.UpdatedFields.NumberSoldSinceUpdate)
}

func NewWeaviate(ctx context.Context) *Weaviate {
	weaviateClient := &Weaviate{}
	weaviateClient.ctx = ctx

	return weaviateClient
}

func (weaviate *Weaviate) Start() chan ItemChange {
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

	weaviate.itemChangeChannel = make(chan ItemChange)
	weaviate.ctx = context.Background()

	return weaviate.itemChangeChannel
}

func (weaviate *Weaviate) DispatchChanges(itemChange ItemChange) {

	if itemChange.Mode == "CREATE" {
		// Create object with correct properties on Weaviate.
		w, err := weaviate.Client.
			Data().
			Creator().
			WithClassName("item").
			WithProperties(
				map[string]interface{}{
					"name":                  itemChange.Item.Name,
					"categoryName":          itemChange.Item.CategoryName,
					"imageURL":              itemChange.Item.Photo,
					"price":                 itemChange.Item.Price,
					"numberSoldSinceUpdate": itemChange.Item.NumberSoldSinceUpdate,
					"dateLastSold":          itemChange.Item.DateLastSold,
				}).
			Do(weaviate.ctx)

		if err != nil {
			panic(err)
		}

		itemChange.Item.WeaviateID = w.Object.ID.String()

	} else if itemChange.Mode == "UPDATE" {
		// Update properties on Weaviate.

		itemUpdates := map[string]interface{}{}

		if itemChange.UpdatedFields.Name {
			itemUpdates["name"] = itemChange.Item.Name
		}

		if itemChange.UpdatedFields.CategoryName {
			itemUpdates["categoryName"] = itemChange.Item.CategoryName
		}

		if itemChange.UpdatedFields.Photo {
			itemUpdates["imageURL"] = itemChange.Item.Photo
		}

		if itemChange.UpdatedFields.Price {
			itemUpdates["price"] = itemChange.Item.Price
		}

		if itemChange.UpdatedFields.NumberSoldSinceUpdate {
			itemUpdates["numberSoldSinceUpdate"] = itemChange.Item.NumberSoldSinceUpdate
		}

		if itemChange.UpdatedFields.DateLastSold {
			itemUpdates["dateLastSold"] = itemChange.Item.DateLastSold
		}

		err := weaviate.Client.
			Data().
			Updater().
			WithMerge().
			WithClassName("item").
			WithProperties(itemUpdates).
			Do(weaviate.ctx)

		if err != nil {
			panic(err)
		}

	} else if itemChange.Mode == "DELETE" {
		// Use the client to delete the item from Weaviate.
		err := weaviate.Client.
			Data().
			Deleter().
			WithClassName("item").
			WithID(itemChange.Item.WeaviateID).
			Do(weaviate.ctx)

		if err != nil {
			panic(err)
		}

	} else {
		// There is an error in the invoking code, we should not receive any other Item Mode.
		log.Fatal("Invalid Item Change Mode.")
	}

}

func (weaviate *Weaviate) ItemChangeHandler(entClient *ent.Client) {

	for itemChange := range weaviate.itemChangeChannel {
		go weaviate.DispatchChanges(itemChange)
	}
}
