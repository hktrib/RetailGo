package weaviate

import (
	"context"
	"log"

	server "github.com/hktrib/RetailGo/cmd/api"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/util"
	weaviateClient "github.com/weaviate/weaviate-go-client/v4/weaviate"
	weaviateAuth "github.com/weaviate/weaviate-go-client/v4/weaviate/auth"
	"github.com/weaviate/weaviate/entities/models"
)

type Weaviate struct {
	Client            *weaviateClient.Client
	ctx               context.Context
	itemChangeChannel chan ItemChange
}

// Define a type for Item-Change Requests (Create, Update, Delete)

type ItemChange struct {
	Item          ent.Item
	Mode          string
	UpdatedFields server.UpdatedFields
}

func vectorUpdateNeeded(itemChange ItemChange) bool {
	return itemChange.Mode == "CREATE" || (itemChange.Mode == "UPDATE" && !(itemChange.UpdatedFields.Name || itemChange.UpdatedFields.CategoryName || itemChange.UpdatedFields.Price))
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

	itemClassObj := &models.Class{
		Class:        "Item",
		Vectorizer:   "none",
		ModuleConfig: map[string]interface{}{},
	}

	err = weaviate.Client.Schema().ClassCreator().WithClass(itemClassObj).Do(context.Background())

	if err != nil {
		panic(err)
	}

	userClassObj := &models.Class{
		Class:        "User",
		Vectorizer:   "none",
		ModuleConfig: map[string]interface{}{},
	}

	err = weaviate.Client.Schema().ClassCreator().WithClass(userClassObj).Do(context.Background())

	if err != nil {
		panic(err)
	}

	weaviate.itemChangeChannel = make(chan ItemChange)

	return weaviate.itemChangeChannel
}

func (weaviate *Weaviate) DispatchChanges(itemChange ItemChange) {

	if itemChange.Mode == "Create" {
		// Update literal details on Weaviate.
		w, err := weaviate.Client.
			Data().
			Creator().
			WithClassName("item").
			WithProperties(
				map[string]interface{}{
					"name":         itemChange.Item.Name,
					"categoryName": itemChange.Item.CategoryName,
					"imageURL":     itemChange.Item.Photo,
					"price":        itemChange.Item.Price,
				}).
			Do(weaviate.ctx)

		if err != nil {
			itemChange.Item.WeaviateID = w.Object.ID.String()
		}

	} else if itemChange.Mode == "Update" {
		// Update literal details on Weaviate.

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

		err := weaviate.Client.
			Data().
			Updater().
			WithMerge().
			WithClassName("item").
			WithProperties(itemUpdates).
			Do(weaviate.ctx)

		// If only unvectorized fields are changed and if done so successfully, exit.
		if err == nil && !(itemChange.UpdatedFields.Name || itemChange.UpdatedFields.CategoryName || itemChange.UpdatedFields.Photo) {
			return
		}

	} else if itemChange.Mode == "Delete" {
		// Use the client to delete the item from Weaviate and exit if successful.
		err := weaviate.Client.
			Data().
			Deleter().
			WithClassName("Item").
			WithID(itemChange.Item.WeaviateID).
			Do(weaviate.ctx)

		if err == nil {
			return
		}

	} else {
		// There is an error in the invoking code, we should not receive any other Item Mode.
		log.Fatal("Invalid Item Change Mode.")
	}

}

func (weaviate *Weaviate) DispatchVector(itemChange ItemChange, entClient *ent.Client) {
	// Mark the items in PG as dirty
	targetItem, err := entClient.Item.Query().Where(item.ID(itemChange.Item.ID)).Only(weaviate.ctx)

	if err != nil || !targetItem.Vectorized {
		return
	}

	targetItem.Update().SetVectorized(false)
}

func (weaviate *Weaviate) ItemChangeHandler(entClient *ent.Client) {

	for itemChange := range weaviate.itemChangeChannel {
		if vectorUpdateNeeded(itemChange) {
			go weaviate.DispatchVector(itemChange, entClient)
		}

		go weaviate.DispatchChanges(itemChange)
	}
}
