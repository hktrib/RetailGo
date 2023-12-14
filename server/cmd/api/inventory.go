package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/categoryitem"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/weaviate"
)

/*
InvRead Brief:

	________________________________________

-> Reads all items from specified store id given in URL (Ex: store/{store_id})

Response:

-> On success, return 200 and a list of items in message body.
-> If invalid store_id is provided, return 400.
*/
func (srv *Server) InvRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		log.Debug().Err(err).Msg("InvRead: invalid store_id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Get all items belonging to the store
	inventory, err := srv.DBClient.Item.Query().
		Where(item.StoreID(store_id)).
		All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("InvCreate: unable to query for store's inventory")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Format and return
	inventory_bytes, err := json.Marshal(PruneItems(inventory...))
	if err != nil {
		log.Debug().Err(err).Msg("InvCreate: unable to marshal items in inventory")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(inventory_bytes)
}

/*
InvCreate Brief:

	________________________________________

-> Creates an item in a specified store's inventory (Ex: store/{store_id}/inventory/create), will create the category if it doesn't exist.

Response:

-> On success, return 201.
-> If an invalid store id or
*/
func (srv *Server) InvCreate(w http.ResponseWriter, r *http.Request) {
	// name, photo, quantity, category name

	ctx := r.Context()

	// Load store_id first

	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Load the message parameters (Name, Photo, Quantity, Category Name)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqItem := ent.Item{}
	reqItem.StoreID = store_id
	body_parse_err := json.Unmarshal(req_body, &reqItem)

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create item in database with name, photo, quantity, store_id, category name.
	// get the stores stripe id
	targetStore, err := srv.DBClient.Store.Query().Where(store.IDEQ(store_id)).Only(ctx)
	if err != nil {
		fmt.Println("Unable to find store:", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Create item in Stripe at the target store.
	ProductId, err := CreateStripeItem(&reqItem, targetStore)
	if err != nil {
		fmt.Println("Stripe Create didn't work:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create Item in Database
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

	if create_err != nil {
		fmt.Println("Create didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = createOrUpdateCategory(reqItem.CategoryName, createdItem, ctx, srv)

	if err != nil {

	}

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
	w.Write([]byte("Created"))

}

/*
	@Params
		categoryName -> Category Name (string)
		createdItem -> a created *ent.Item
		ctx -> the request context
		srv -> a server object

	Creates category if the category name is blank,
*/

func createOrUpdateCategory(categoryName string, createdItem *ent.Item, ctx context.Context, srv *Server) error {
	var cat *ent.Category
	var err error

	// If no categoryName provided, find the default "Uncategorized" category for this store. Otherwise, query for the category by name under the store.
	if categoryName == "" {
		cat, err = srv.DBClient.Category.Query().Where(category.Name("Uncategorized"), category.StoreID(createdItem.StoreID)).Only(ctx)
	} else {
		cat, err = srv.DBClient.Category.Query().Where(category.Name(categoryName), category.StoreID(createdItem.StoreID)).Only(ctx)
	}

	if err != nil {
		// If the category is not found, create the category and the item within it.
		if ent.IsNotFound(err) {
			return CreateCatWithItem(createdItem, categoryName, ctx, srv)
		}

		log.Debug().Err(err).Msgf("Failed to query %s", categoryName)
		return err
	}

	// Otherwise, add the createdItem to the category.
	_, err = srv.DBClient.Category.UpdateOne(cat).AddItems(createdItem).Save(ctx)
	return err
}
func swapCategories(item1 *ent.Item, newCatName string, ctx context.Context, srv *Server) error {
	oldCat, err := srv.DBClient.Category.Query().Where(category.Name(item1.CategoryName), category.HasStoreWith(store.ID(item1.StoreID))).Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to query old category")
		return err
	}

	newCat, err := srv.DBClient.Category.Query().Where(category.Name(newCatName)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return CreateCatWithItem(item1, newCatName, ctx, srv)
		}
		log.Debug().Err(err).Msg("Failed to query new category")
		return err
	}

	_, err = srv.DBClient.Category.UpdateOne(oldCat).RemoveItems(item1).Save(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to remove item from old category")
		return err
	}

	_, err = srv.DBClient.Category.UpdateOne(newCat).AddItems(item1).Save(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Failed to add item to new category")
		return err
	}

	return nil
}

func CreateCatWithItem(item1 *ent.Item, newCatName string, ctx context.Context, srv *Server) error {
	_, err := srv.DBClient.Category.Create().
		SetName(newCatName).
		SetStoreID(item1.StoreID).
		AddItems(item1).
		SetPhoto([]byte("hello")).Save(ctx)

	if err != nil {
		log.Debug().Err(err).Msgf("Failed to create %s", newCatName)
		return err
	}
	return nil
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
		log.Debug().Err(err).Msg("InvUpdate: server failed to read message body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reqItem := ent.Item{}

	body_parse_err := json.Unmarshal(req_body, &reqItem)

	// If any are not present or not do not meet requirements (type for example), BadRequest

	if body_parse_err != nil {
		log.Debug().Err(err).Msg("InvUpdate: unmarshalling failed")
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
	if targetItem.CategoryName != reqItem.CategoryName {
		err := swapCategories(targetItem, reqItem.CategoryName, r.Context(), srv)
		if err != nil {
			log.Debug().Err(err).Msg("InvUpdate: failed to swap categories")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	targetStore, err := srv.DBClient.Store.Query().Where(store.IDEQ(targetItem.StoreID)).Only(r.Context())
	if err != nil {
		fmt.Println("Unable to find store:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if targetItem.Name != reqItem.Name {
		np, err := UpdateStripeItem(targetItem, reqItem.Name, targetStore.StripeAccountID)
		if err != nil {
			log.Debug().Err(err).Msg("InvUpdate: failed to update stripe item")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		reqItem.StripeProductID = np.ID
		reqItem.StripePriceID = np.DefaultPrice.ID
	}
	if targetItem.Price != reqItem.Price {
		// get the stores stripe id

		p, err := UpdateStripePrice(targetItem, reqItem.Price, targetStore.StripeAccountID)
		if err != nil {
			log.Debug().Err(err).Msg("InvUpdate: failed to update stripe price")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		reqItem.StripePriceID = p.ID
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
		SetNumberSoldSinceUpdate(originalItem.NumberSoldSinceUpdate).
		Save(r.Context())
	if err != nil {
		log.Debug().Err(err).Msg("InvUpdate: unable to update item in database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	weaviateUpdateErr := srv.WeaviateClient.EditItem(targetItem, fieldsUpdated)
	if weaviateUpdateErr != nil {
		// Rolls back Database update if Weaviate update fails, but no guarantee of rollback working..
		log.Debug().Err(weaviateUpdateErr).Msg("InvUpdate: weaviate failed to update item")
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
	w.Write([]byte("OK"))
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
		log.Debug().Err(err).Msg("InvDelete: unable to parse store_id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	itemId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Debug().Err(err).Msg("InvDelete: unable to get id from URL params")
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
		log.Debug().Err(err).Msg("InvDelete: server failed to find target item")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	weaviateID := targetItem.WeaviateID
	err = DetangleItem(srv.DBClient, targetItem)
	if err != nil {
		log.Debug().Err(err).Msg("InvDelete: server failed to detangle item")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = srv.DBClient.
		Item.DeleteOneID(itemId).
		Where(item.StoreID(store_id)).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		log.Debug().Err(err).Msg("InvDelete: server failed to find target item")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		log.Debug().Err(err).Msg("InvDelete: server failed to delete target item")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = srv.WeaviateClient.DeleteItem(weaviateID)

	if err != nil {
		log.Debug().Err(err).Msg("InvDelete: weaviate failed to delete target item")
		// Should Roll back delete - for now, naively, by just creating a new item -, but doesn't do this yet.
		// _, _ = srv.DBClient.Item.Create().SetQuantity(originalItem.Quantity).SetPhoto(originalItem.Photo).SetName(originalItem.Name).SetPrice(originalItem.Price).SetCategoryName(originalItem.CategoryName).SetVectorized(originalItem.Vectorized).SetNumberSoldSinceUpdate(originalItem.NumberSoldSinceUpdate).SetDateLastSold(originalItem.DateLastSold).SetStore(originalItem.Edges.Store).Set.Save(r.Context())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (srv *Server) InvUpdatePhoto(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	type photoUpdate struct {
		ID    int    `json:"id"`
		Photo string `json:"photo"`
	}

	var req photoUpdate

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := srv.DBClient.Item.UpdateOneID(req.ID).
		Where().
		SetPhoto(req.Photo).
		Save(ctx)

	if err != nil {
		log.Debug().Err(err).Msg("InvUpdatePhoto: Failed to update photo")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func DetangleItem(Client *ent.Client, item *ent.Item) error {
	// Get the item's category

	// Remove the item from the category
	_, err := Client.CategoryItem.Delete().Where(categoryitem.ItemID(item.ID)).Exec(context.Background())
	if err != nil {
		log.Debug().Err(err).Msg("DetangleItem: failed to remove item from category")
		return err
	}

	return nil
}
