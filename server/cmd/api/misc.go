package server

import (
	"encoding/json"

	"github.com/hktrib/RetailGo/internal/ent"
)

// Item with only the feilds the client should see
type ClientItem ent.Item
type ClientStore ent.Store
type ClientUserStore ent.UserToStore

// Overload for default json marshaller
func (i ClientItem) MarshalJSON() ([]byte, error) {
	return MarshalItem(i)
}

func (i ClientStore) MarshalJSON() ([]byte, error) {
	return MarshalStore(i)
}
func (i ClientUserStore) MarshalJSON() ([]byte, error) {
	return MarshelUserStore(i)
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

// Overload for default json marshaller
func MarshalStore(TargetItem ClientStore) ([]byte, error) {
	/*
		cats, err := temp.Category.Query().Where(category.HasItemsWith(item.ID(TargetItem.ID))).All(context.Background())
		var catNames []string
		if err != nil {
			if ent.IsNotFound(err) {
				catNames[0] = "Uncategorized"
			} else {
				return nil, err
			}
		}
		for _, cat := range cats {
			catNames = append(catNames, cat.Name)
		}

	*/
	return json.Marshal(map[string]interface{}{
		"id":         TargetItem.ID,
		"uuid":       TargetItem.UUID,
		"store_name": TargetItem.StoreName,
	})
}
func PruneStore(TargetItem *ent.Store) ClientStore {
	var store ClientStore
	store = ClientStore(*TargetItem)
	return store

}

// Overload for default json marshaller
func MarshelUserStore(TargetItem ClientUserStore) ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"user_id":          TargetItem.UserID,
		"store_id":         TargetItem.StoreID,
		"permission_level": TargetItem.PermissionLevel,
		"store_name":       TargetItem.StoreName,
	})
}
func PruneUserStore(TargetItems ...*ent.UserToStore) []ClientUserStore {
	var clientItems []ClientUserStore

	for _, item := range TargetItems {
		clientItems = append(clientItems, ClientUserStore(*item))
	}
	if clientItems == nil {
		clientItems = make([]ClientUserStore, 0)
	}
	return clientItems

}
