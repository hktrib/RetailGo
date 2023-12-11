package server

import (
	"encoding/json"

	"github.com/hktrib/RetailGo/internal/ent"
)

// Item with only the feilds the client should see
type ClientItem ent.Item

// Overload for default json marshaller
func (i ClientItem) MarshalJSON() ([]byte, error) {
	return MarshalItem(i)
}

// Overload for default json marshaller
func MarshalItem(TargetItem ClientItem) ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"id":                TargetItem.ID,
		"name":              TargetItem.Name,
		"photo":             TargetItem.Photo,
		"quantity":          TargetItem.Quantity,
		"price":             TargetItem.Price,
		"category_name":     TargetItem.CategoryName,
		"stripe_product_id": TargetItem.StripeProductID,
		"stripe_price_id":   TargetItem.StripePriceID,
	})
}
func PruneItems(TargetItems ...*ent.Item) []ClientItem {
	var clientItems []ClientItem

	for _, item := range TargetItems {
		clientItems = append(clientItems, ClientItem(*item))
	}
	if clientItems == nil {
		clientItems = make([]ClientItem, 0)
	}
	return clientItems

}
