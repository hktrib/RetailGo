package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/weaviate"
)

/*
InvRead Brief:

	________________________________________

-> Read's all items from Specified Store's (Ex: store/{store_id}) Inventory

External Package Calls:

	N/A
*/
func (srv *Server) InvRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("Error with Store Id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Get all items belonging to the store
	inventory, err := srv.DBClient.Item.Query().
		Where(item.StoreID(store_id)).
		All(ctx)
	if err != nil {
		fmt.Println("Issue with querying the DB.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Format and return
	inventory_bytes, err := json.Marshal(PruneItems(inventory...))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(inventory_bytes)
}

/*
InvCreate Brief:

	________________________________________

-> Creates an item in a specified store's inventory (Ex: store/{store_id}/inventory/create)

External Package Calls:

	N/A
*/
func (srv *Server) InvCreate(w http.ResponseWriter, r *http.Request) {
	// name, photo, quantity, store_id, category

	ctx := r.Context()

	// Load store_id first

	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqItem := ent.Item{}

	body_parse_err := json.Unmarshal(req_body, &reqItem)

	// If any are not present or not do not meet requirements (type for example), BadRequest

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Create item in database with name, photo, quantity, store_id, category

	ProductId, err := CreateStripeItem(&reqItem)
	if err != nil {
		fmt.Println("Stripe Create didn't work:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// set all item parameters
	createdItem, create_err := srv.DBClient.
		Item.Create().
		SetPrice(float64(reqItem.Price)).
		SetName(reqItem.Name).
		SetStripePriceID(ProductId.DefaultPrice.ID).
		SetStripeProductID(ProductId.ID).
		SetPhoto(reqItem.Photo).
		SetQuantity(reqItem.Quantity).
		SetStoreID(store_id).
		SetCategoryName(reqItem.CategoryName).
		SetWeaviateID("").
		SetVectorized(false).
		SetNumberSoldSinceUpdate(0).
		SetDateLastSold("").
		Save(ctx)
	// If this create doesn't work, InternalServerError
	if create_err != nil {
		fmt.Println("Create didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": createdItem.ID,
	})

	weaviateID, weaviateErr := srv.WeaviateClient.CreateItem(createdItem)

	fmt.Println("Weaviate ID on create:", weaviateID)

	// Currently raises 500 if creation of weaviate copy or storing the weaviate id fails.
	// Attempts to roll back the database in response to this so that there are no ghost items for Weaviate.

	if weaviateErr != nil {
		fmt.Println("Weaviate Create didn't work")
		_ = srv.DBClient.
			Item.DeleteOneID(createdItem.ID).
			Where(item.StoreID(store_id)).
			Exec(r.Context())

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, setWeaviateIDErr := createdItem.Update().SetWeaviateID(weaviateID).Save(r.Context())

	if setWeaviateIDErr != nil {
		fmt.Println("Failed to set Weaviate ID")
		_ = srv.DBClient.
			Item.DeleteOneID(createdItem.ID).
			Where(item.StoreID(store_id)).
			Exec(r.Context())

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)

}

/*
InvUpdate Brief:

	________________________________________

-> Updates an item in a specified store's inventory (Ex: store/{store_id}/inventory/update)

External Package Calls:

	N/A
*/
func (srv *Server) InvUpdate(w http.ResponseWriter, r *http.Request) {
	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqItem := ent.Item{}

	body_parse_err := json.Unmarshal(req_body, &reqItem)

	// If any are not present or not do not meet requirements (type for example), BadRequest

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetItem, err := srv.DBClient.Item.
		Query().
		Where(item.ID(reqItem.ID)).
		Only(r.Context())

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	originalItem := ent.Item{
		Name:         targetItem.Name,
		Photo:        targetItem.Photo,
		Quantity:     targetItem.Quantity,
		Price:        targetItem.Price,
		CategoryName: targetItem.CategoryName,
		Vectorized:   targetItem.Vectorized,
	}

	log.Debug().Msg(fmt.Sprintf("originalItem: %v \n requested item: %v", originalItem, reqItem))

	fieldsUpdated := weaviate.UpdatedFields{
		Name:                  targetItem.Name != reqItem.Name,
		Photo:                 targetItem.Photo != reqItem.Photo,
		Quantity:              targetItem.Quantity != reqItem.Quantity,
		Price:                 targetItem.Price != reqItem.Price,
		CategoryName:          targetItem.CategoryName != reqItem.CategoryName,
		NumberSoldSinceUpdate: false, // Assume that sales is taken care of elsewhere. If not, this is subject to change.
		DateLastSold:          false, // Assume that sales is taken care of elsewhere. If not, this is subject to change.
	}

	// Update the item's paramaters in the database
	_, err = targetItem.Update().
		SetQuantity(reqItem.Quantity).
		SetName(reqItem.Name).
		SetPhoto(reqItem.Photo).
		SetPrice(reqItem.Price).
		SetCategoryName(reqItem.CategoryName).
		SetVectorized(originalItem.Vectorized && !(weaviate.ChangesVectorizedProperties(fieldsUpdated))).
		Save(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	weaviateUpdateErr := srv.WeaviateClient.EditItem(targetItem, fieldsUpdated)
	if weaviateUpdateErr != nil {
		// Rolls back Database update if Weaviate update fails, but no guarantee of rollback working..
		log.Debug().Err(weaviateUpdateErr).Msg("Failed to edit item in Weaviate")
		_, _ = targetItem.Update().
			SetQuantity(originalItem.Quantity).
			SetPhoto(originalItem.Photo).
			SetName(originalItem.Name).
			SetPrice(originalItem.Price).
			SetCategoryName(originalItem.CategoryName).
			SetVectorized(originalItem.Vectorized).
			Save(r.Context())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

/*
InvDelete Brief:

	________________________________________

-> Deletes an item in a specified store's inventory (Ex: store/{store_id}/inventory/update)

External Package Calls:

	N/A
*/
func (srv *Server) InvDelete(w http.ResponseWriter, r *http.Request) {

	// store := r.Context().Value("store_var").(*ent.Store)

	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Store Id parsing failed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	itemId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Debug().Err(err).Msg("InvDelete")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetItem, err := srv.DBClient.Item.
		Query().
		Where(item.ID(itemId)).
		Only(r.Context())

	// originalItem := ent.Item{
	// 	Name:         targetItem.Name,
	// 	Photo:        targetItem.Photo,
	// 	Quantity:     targetItem.Quantity,
	// 	Price:        targetItem.Price,
	// 	CategoryName: targetItem.CategoryName,
	// 	Vectorized:   targetItem.Vectorized,
	// }

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	weaviateID := targetItem.WeaviateID

	err = srv.DBClient.
		Item.DeleteOneID(itemId).
		Where(item.StoreID(store_id)).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = srv.WeaviateClient.DeleteItem(weaviateID)

	if err != nil {
		log.Debug().Err(err).Msg("Delete in Weaviate Failed")
		// Should Roll back delete - for now, naively, by just creating a new item -, but doesn't do this yet.
		// _, _ = srv.DBClient.Item.Create().SetQuantity(originalItem.Quantity).SetPhoto(originalItem.Photo).SetName(originalItem.Name).SetPrice(originalItem.Price).SetCategoryName(originalItem.CategoryName).SetVectorized(originalItem.Vectorized).SetNumberSoldSinceUpdate(originalItem.NumberSoldSinceUpdate).SetDateLastSold(originalItem.DateLastSold).SetStore(originalItem.Edges.Store).Set.Save(r.Context())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
