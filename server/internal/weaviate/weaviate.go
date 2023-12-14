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

func ChangesVectorizedProperties(originalItem *ent.Item, updatedItem *ent.Item) bool {
	updatedFields := UpdatedFields{
		Name:                  originalItem.Name != updatedItem.Name,
		Photo:                 originalItem.Photo != updatedItem.Photo,
		Quantity:              originalItem.Quantity != updatedItem.Quantity,
		Price:                 originalItem.Price != updatedItem.Price,
		CategoryName:          originalItem.CategoryName != updatedItem.CategoryName,
		NumberSoldSinceUpdate: false, // Assume that sales is taken care of elsewhere. If not, this is subject to change.
		DateLastSold:          false, // Assume that sales is taken care of elsewhere. If not, this is subject to change.
	}

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
