package weaviate

import (
	"context"

	weaviateClient "github.com/weaviate/weaviate-go-client/v4/weaviate"
	weaviateAuth "github.com/weaviate/weaviate-go-client/v4/weaviate/auth"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/util"
)

type Weaviate struct {
	Client *weaviateClient.Client
	ctx    context.Context
}

func ChangesVectorizedProperties(updatedFields UpdatedFields) bool {
	return (updatedFields.Name || updatedFields.CategoryName || updatedFields.Price || updatedFields.Photo || updatedFields.DateLastSold || updatedFields.NumberSoldSinceUpdate)
}

func NewWeaviate(ctx context.Context) *Weaviate {
	weaviateClient := &Weaviate{}
	weaviateClient.ctx = ctx

	return weaviateClient
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

	weaviate.ctx = context.Background()
}

func (weaviate *Weaviate) CreateItem(item *ent.Item) (string, error) {
	w, err := weaviate.Client.
		Data().
		Creator().
		WithClassName("item").
		WithProperties(
			map[string]interface{}{
				"name":                  item.Name,
				"categoryName":          item.CategoryName,
				"imageURL":              item.Photo,
				"price":                 item.Price,
				"numberSoldSinceUpdate": item.NumberSoldSinceUpdate,
				"dateLastSold":          item.DateLastSold,
			}).
		Do(weaviate.ctx)

	if err != nil {
		return "", err
	}

	return w.Object.ID.String(), err
}

func (weaviate *Weaviate) EditItem(item *ent.Item, updatedFields UpdatedFields) error {
	// Update properties on Weaviate.

	itemUpdates := map[string]interface{}{}

	if updatedFields.Name {
		itemUpdates["name"] = item.Name
	}

	if updatedFields.CategoryName {
		itemUpdates["categoryName"] = item.CategoryName
	}

	if updatedFields.Photo {
		itemUpdates["imageURL"] = item.Photo
	}

	if updatedFields.Price {
		itemUpdates["price"] = item.Price
	}

	if updatedFields.NumberSoldSinceUpdate {
		itemUpdates["numberSoldSinceUpdate"] = item.NumberSoldSinceUpdate
	}

	if updatedFields.DateLastSold {
		itemUpdates["dateLastSold"] = item.DateLastSold
	}

	err := weaviate.Client.
		Data().
		Updater().
		WithID(item.WeaviateID).
		WithClassName("item").
		WithProperties(itemUpdates).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	return err
}

func (weaviate *Weaviate) DeleteItem(weaviateID string) error {
	err := weaviate.Client.
		Data().
		Deleter().
		WithClassName("item").
		WithID(weaviateID).
		Do(weaviate.ctx)

	return err
}
