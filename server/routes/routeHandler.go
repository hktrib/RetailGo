package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/ent/item"
	store2 "github.com/hktrib/RetailGo/ent/store"
	"github.com/hktrib/RetailGo/util"
	"net/http"
	"strconv"
)

type RouteHandler struct {
	Client *ent.Client
	Config *util.Config
}

func (rh *RouteHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World -> RetailGo!!"))

	ctx := r.Context()

	user, err := rh.Client.User.
		Create().
		SetUsername("gvadhul").
		SetRealName("Giridhar Vadhul").
		SetEmail("gvadhul@ucsc.edu").
		SetIsOwner(true).
		SetStoreID(1381).
		Save(ctx)

	if err != nil {
		fmt.Println("Create user failed")
		fmt.Println(err)
	}

	store, err := rh.Client.Store.
		Create().
		SetStoreName("Giridhar's Test Store").
		SetID(1391).
		Save(ctx)

	if err != nil {
		fmt.Println("Create Store Failed")
		fmt.Println(err)
	}

	item1, err := rh.Client.Item.Create().SetName("Giridhar Food 1").SetStoreID(1391).SetPhoto([]byte("myimage")).SetQuantity(1).SetCategory("Food").Save(ctx)
	if err != nil {

		fmt.Println(err)
	}
	item2, err := rh.Client.Item.Create().SetName("Giridhar Food 2").SetStoreID(1391).SetPhoto([]byte("anotherimage")).SetQuantity(2).SetCategory("Grocery").Save(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(store)
	fmt.Printf("Item1: %s\n", item1)
	fmt.Printf("Item2: %s\n", item2)
}

// API to create a new inventory item

func (rh *RouteHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	// Take an item's:
	// Name
	// Photo
	// Quantity
	// Store ID
	// Category

}

// API to read inventory of a store
// Route: inventory/?store_id=_

func (rh *RouteHandler) ReadAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {

		log.Fatal("Not a valid store ID")
	}

	// Verify valid store id (not SQL Injection)

	// Run the SQL Query by store id on the items table, to get all items belonging to this store
	inventory, err := rh.Client.Item.Query().Where(item.StoreID(store_id)).All(ctx)

	if err != nil {
		fmt.Printf("Failed to retrieve items for this store: %s", err)
	}

	inventory_bytes, err := json.MarshalIndent(inventory, "", "")

	// Format and return

	if err != nil {
		fmt.Println("Failed to Convert Inventory to JSON properly")
	} else {
		w.Write(inventory_bytes)
	}
}

func (rh *RouteHandler) ValidateStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store_id := chi.URLParam(r, "store_id")
		StoreId, err := strconv.Atoi(store_id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		store, err := rh.Client.Store.
			Query().Where(store2.ID(StoreId)).Only(r.Context())
		if ent.IsNotFound(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "store_var", store)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (rh *RouteHandler) InvUpdate(w http.ResponseWriter, r *http.Request) {
	// get item id from url query string

	itemId, err := strconv.Atoi(r.URL.Query().Get("item"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	change, err := strconv.Atoi(r.URL.Query().Get("change"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetItem, err := rh.Client.Item.
		Query().
		Where(item.ID(itemId)).
		Only(r.Context())
	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = targetItem.Update().SetQuantity(targetItem.Quantity + change).
		Save(r.Context())

	if err != nil {
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
func (rh *RouteHandler) InvDelete(w http.ResponseWriter, r *http.Request) {

	store := r.Context().Value("store_var").(*ent.Store)

	item_id, err := strconv.Atoi(r.URL.Query().Get("item"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = rh.Client.
		Item.DeleteOneID(item_id).
		Where(item.StoreID(store.ID)).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))

}
