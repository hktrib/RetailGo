package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/ent/item"
	"github.com/hktrib/RetailGo/util"
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
	fmt.Println("User:")
	fmt.Println(user)
	fmt.Println("Store:")
	fmt.Println(store)
	fmt.Printf("Item1: %s\n", item1)
	fmt.Printf("Item2: %s\n", item2)
}

// API to create a new inventory item

func (rh *RouteHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
}

// API to read inventory of a store

func (rh *RouteHandler) ReadAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(r.URL.Query().Get("store_id"))

	if err != nil {
		log.Fatal("Not a valid store ID")
	}

	// Verify valid store id (not SQL Injection)

	// Run the SQL Query by store id on the items table, to get all items belonging to this store
	inventory, err := rh.Client.Item.Query().Where(item.StoreID(store_id)).All(ctx)

	if err != nil {
		fmt.Printf("Failed to retrieve items for this store: %s", err)
	}

	for _, item := range inventory {
		fmt.Printf("Item:, %s", item)
	}

	// Format and return

	w.Write([]byte("You have reached ReadAll.\n"))
}
