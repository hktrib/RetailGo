package server

import (
	"encoding/json"
	"github.com/hktrib/RetailGo/internal/ent"
)

// ClientItem item that is sent to the client
type ClientItem ent.Item

func (ite ClientItem) MarshalJSON() ([]byte, error) {
	return MarshalItem(ite)
}

func MarshalItem(TargetItem ClientItem) ([]byte, error) {
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
		"id":       TargetItem.ID,
		"name":     TargetItem.Name,
		"photo":    TargetItem.Photo,
		"quantity": TargetItem.Quantity,
		"price":    TargetItem.Price,
		"category": TargetItem.CategoryName,
	})
}
func PruneItems(TargetItems ...*ent.Item) []ClientItem {
	var clientItems []ClientItem
	for _, item := range TargetItems {
		clientItems = append(clientItems, ClientItem(*item))
	}
	return clientItems

}
