package server

import (
	"encoding/json"
	"fmt"
	"io"

	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
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

	user, err := srv.DBClient.User.
		Create().
		SetFirstName("Giridhar").
		SetLastName("Vadhul").
		SetEmail("gvadhul@ucsc.edu").
		SetIsOwner(true).
		Save(ctx)

	if err != nil {
		fmt.Println("Create user failed")
		fmt.Println(err)
	}

	store, err := srv.DBClient.Store.
		Create().
		SetStoreName("Giridhar's Test Store").
		SetID(1391).
		Save(ctx)

	if err != nil {
		fmt.Println("Create Store Failed")
		fmt.Println(err)
	}

	item1, err := srv.DBClient.Item.Create().SetName("Giridhar Food 1").SetStoreID(1391).SetPhoto([]byte("myimage")).SetQuantity(1).Save(ctx)
	if err != nil {

		fmt.Println(err)
	}
	item2, err := srv.DBClient.Item.Create().SetName("Giridhar Food 2").SetStoreID(1391).SetPhoto([]byte("anotherimage")).SetQuantity(2).Save(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(store)
	fmt.Printf("Item1: %s\n", item1)
	fmt.Printf("Item2: %s\n", item2)
}

func (srv *Server) InvRead(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Error with Store Id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Verify valid store id (not SQL Injection)

	// Run the SQL Query by store id on the items table, to get all items belonging to this store
	inventory, err := srv.DBClient.Item.Query().Where(item.StoreID(store_id)).All(ctx)

	if err != nil {
		fmt.Println("Issue with querying the DB.")
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
	}

	fmt.Println("reqItem:", reqItem.Price, reqItem.Name, ProductId.DefaultPrice.ID, reqItem.Photo, reqItem.Quantity, store_id)

	createdItem, create_err := srv.DBClient.Item.Create().SetPrice(float64(reqItem.Price)).SetName(reqItem.Name).SetStripePriceID(ProductId.DefaultPrice.ID).SetStripeProductID(ProductId.ID).SetPhoto([]byte(reqItem.Photo)).SetQuantity(reqItem.Quantity).SetStoreID(store_id).SetCategoryName(reqItem.CategoryName).Save(ctx)
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

	// ctx := r.Context()

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

	_, err = targetItem.Update().SetQuantity(reqItem.Quantity).SetName(reqItem.Name).SetPhoto(reqItem.Photo).SetPrice(reqItem.Price).SetCategoryName(reqItem.CategoryName).Save(r.Context())

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

	// store := r.Context().Value("store_var").(*ent.Store)

	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Store Id parsing failed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item_id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println("Item id parsing failed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.
		Item.DeleteOneID(item_id).
		Where(item.StoreID(store_id)).
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
