package webhook

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/weaviate"
)

type ItemChange struct {
	Type      string
	Table     string
	Schema    string
	Record    ent.Item
	OldRecord ent.Item
}

func parseItemChange(w http.ResponseWriter, r *http.Request) *ItemChange {
	defer r.Body.Close()

	payload, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error Reading Body:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil
	}

	fmt.Printf("Logging Request: %s", payload)

	var itemChange *ItemChange = &ItemChange{}

	err = json.Unmarshal(payload, itemChange)

	if err != nil {
		fmt.Println("Error Unmarshalling:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return nil
	}

	return itemChange
}

func validateItemChange(w http.ResponseWriter, itemChange *ItemChange, changeType string) {
	if itemChange.Type != changeType || itemChange.Table != "items" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func findUpdatedFields(updatedItem *ent.Item, oldItem *ent.Item) *weaviate.UpdatedFields {
	var updatedFields *weaviate.UpdatedFields
	updatedFields.Name = updatedItem.Name == oldItem.Name
	updatedFields.Photo = updatedItem.Photo == oldItem.Photo
	updatedFields.Quantity = updatedItem.Quantity == oldItem.Quantity
	updatedFields.Price = updatedItem.Price == oldItem.Price
	updatedFields.NumberSoldSinceUpdate = updatedItem.NumberSoldSinceUpdate == oldItem.NumberSoldSinceUpdate
	updatedFields.DateLastSold = updatedItem.DateLastSold == oldItem.DateLastSold

	return updatedFields
}

func HandleWeaviateCreate(w http.ResponseWriter, r *http.Request, DBClient *ent.Client, WeaviateClient *weaviate.Weaviate) {

	itemChange := parseItemChange(w, r)

	fmt.Println("Parsed Item Change:", itemChange)

	if itemChange == nil {
		return
	}

	validateItemChange(w, itemChange, "INSERT")

	var createdItem *ent.Item = &itemChange.Record

	fmt.Println("createdItem:", createdItem)

	if createdItem == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	weaviateID, weaviateErr := WeaviateClient.CreateItem(createdItem)

	fmt.Println("Weaviate ID on create:", weaviateID)

	// Currently raises 500 if creation of weaviate copy or storing the weaviate id fails.
	// Previously attempted to rollback database changes, Supabase should re-call on error instead now, so this is not necessary.

	if weaviateErr != nil {
		fmt.Println("Weaviate Create didn't work:", weaviateErr)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	createdItem, err := DBClient.Item.Query().Where(item.ID(createdItem.ID)).Only(r.Context()) //Save(r.Context())

	if err != nil {
		fmt.Println("Item did not exist in Supabase yet.", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, setWeaviateIDErr := createdItem.Update().SetWeaviateID(weaviateID).Save(r.Context())

	if setWeaviateIDErr != nil {
		fmt.Println("Failed to set Weaviate ID")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func HandleWeaviateUpdate(w http.ResponseWriter, r *http.Request, DBClient *ent.Client, WeaviateClient *weaviate.Weaviate) {
	itemChange := parseItemChange(w, r)

	if itemChange == nil {
		return
	}

	validateItemChange(w, itemChange, "UPDATE")

	var updatedItem *ent.Item = &itemChange.Record
	var originalItem *ent.Item = &itemChange.OldRecord

	if updatedItem == nil || originalItem == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedFields := findUpdatedFields(updatedItem, originalItem)

	weaviateErr := WeaviateClient.EditItem(updatedItem, updatedFields)

	// Currently raises 500 if update on weaviate copy fails.
	// Previously attempted to rollback database changes, Supabase should re-call on error instead now, so this is not necessary.

	if weaviateErr != nil {
		fmt.Println("Weaviate Update didn't work")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func HandleWeaviateDelete(w http.ResponseWriter, r *http.Request, DBClient *ent.Client, WeaviateClient *weaviate.Weaviate) {
	itemChange := parseItemChange(w, r)

	if itemChange != nil {
		return
	}

	validateItemChange(w, itemChange, "DELETE")

	var deletedItem *ent.Item = &itemChange.OldRecord

	if deletedItem == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if deletedItem.WeaviateID == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}

	err := WeaviateClient.DeleteItem(deletedItem.WeaviateID)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
