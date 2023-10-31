package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/ent/item"
	store2 "github.com/hktrib/RetailGo/ent/store"
)

type TempItem struct {
	Name     string
	Photo    string
	Quantity int
	Category string
	Price    float64
}

func (srv *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World -> RetailGo!!"))

	ctx := r.Context()

	user, err := srv.Client.User.
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

	store, err := srv.Client.Store.
		Create().
		SetStoreName("Giridhar's Test Store").
		SetID(1391).
		Save(ctx)

	if err != nil {
		fmt.Println("Create Store Failed")
		fmt.Println(err)
	}

	item1, err := srv.Client.Item.Create().SetName("Giridhar Food 1").SetStoreID(1391).SetPhoto([]byte("myimage")).SetQuantity(1).SetCategory("Food").Save(ctx)
	if err != nil {

		fmt.Println(err)
	}
	item2, err := srv.Client.Item.Create().SetName("Giridhar Food 2").SetStoreID(1391).SetPhoto([]byte("anotherimage")).SetQuantity(2).SetCategory("Grocery").Save(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(store)
	fmt.Printf("Item1: %s\n", item1)
	fmt.Printf("Item2: %s\n", item2)
}

func (srv *Server) ValidateStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store_id := chi.URLParam(r, "store_id")
		StoreId, err := strconv.Atoi(store_id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		store, err := srv.Client.Store.
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

func (srv *Server) InvRead(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verify valid store id (not SQL Injection)

	// Run the SQL Query by store id on the items table, to get all items belonging to this store
	inventory, err := srv.Client.Item.Query().Where(item.StoreID(store_id)).All(ctx)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	inventory_bytes, err := json.MarshalIndent(inventory, "", "")

	// Format and return

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Write(inventory_bytes)
}

// POST without id (because item needs to be created)
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

	req_item := ent.Item{}

	body_parse_err := json.Unmarshal(req_body, &req_item)

	// If any are not present or not do not meet requirements (type for example), BadRequest

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Create item in database with name, photo, quantity, store_id, category

	createdItem, create_err := srv.Client.Item.Create().SetPrice(float64(req_item.Price)).SetName(req_item.Name).SetPhoto([]byte(req_item.Photo)).SetCategory(req_item.Category).SetQuantity(req_item.Quantity).SetStoreID(store_id).Save(ctx)
	// If this create doesn't work, InternalServerError

	if create_err != nil {
		fmt.Println("Create didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": createdItem.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func (srv *Server) InvUpdate(w http.ResponseWriter, r *http.Request) {
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

	targetItem, err := srv.Client.Item.
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

	_, err = targetItem.Update().SetQuantity(change).
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
func (srv *Server) InvDelete(w http.ResponseWriter, r *http.Request) {

	store := r.Context().Value("store_var").(*ent.Store)

	item_id, err := strconv.Atoi(r.URL.Query().Get("item"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = srv.Client.
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
	w.Write([]byte("OK"))
}
