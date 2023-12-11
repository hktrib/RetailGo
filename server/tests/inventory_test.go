package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	server "github.com/hktrib/RetailGo/cmd/api"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/util"
)

var idVal int

func TestInvRead(tst *testing.T) {
	srv := setupServer()

	/** INCORRECT REQUESTS **/

	req, _ := http.NewRequest("GET", "/store/1231231/inventory", nil)
	rr := executeRequest(req, srv)
	checkResponseCode(tst, http.StatusNotFound, rr.Code)

	/** CORRECT REQUESTS **/

	req, _ = http.NewRequest("GET", "/store/1391/inventory/", nil)
	rr = executeRequest(req, srv)
	checkResponseCode(tst, http.StatusAccepted, rr.Code)

	// Checking Response Body format
	var inventory_items []ent.Item

	err := json.NewDecoder(rr.Body).Decode(&inventory_items)
	if err != nil {
		tst.Logf("Response: %v", inventory_items)
		tst.Error("Error: InvRead Response not in correct format")
	}
}

func TestInvCreate(tst *testing.T) {
	srv := setupServer()

	invItem := ent.Item{
		Name:     "Newest Gerber Grapes",
		Photo:    "uhsrgouhsrgouhsrgoushrg",
		Quantity: 15,
		Price:    120.9,
		StoreID:  141,
	}

	invItemBytes, err := json.Marshal(invItem)
	bytes := bytes.NewBuffer(invItemBytes)
	if err != nil {
		log.Fatal(err)
		tst.Error(err)
	}

	req, _ := http.NewRequest("POST", "/store/141/inventory/create", bytes)

	rr := executeRequest(req, srv)
	checkResponseCode(tst, http.StatusCreated, rr.Code)

	type Item struct {
		ID int `json:"id"`
	}
	it := Item{}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		tst.Error(err)
	}

	err = json.Unmarshal(body, &it)
	if err != nil {
		tst.Error(err)
	}

	idVal = it.ID

	checkInvItem, err := srv.DBClient.Item.Query().Where(item.ID(idVal)).First(req.Context())
	if err != nil {
		log.Fatal(err)
		tst.Error(err)
	}
	if checkInvItem.Name != "Newest Gerber Grapes" {
		tst.Error("Error: Item Created is not Item retrieved")
		tst.Logf("Item Name: %v, Item ID: %v", checkInvItem.Name, checkInvItem.ID)
	}
}

func TestInvItemQuantUpdate(tst *testing.T) {
	srv := setupServer()

	req, _ := http.NewRequest("POST", fmt.Sprintf("/store/1391/inventory/update?item=%v&change=10420690", idVal), nil)

	rr := executeRequest(req, srv)

	checkResponseCode(tst, http.StatusOK, rr.Code)

	checkInvItem, err := srv.DBClient.Item.Query().Where(item.ID(idVal)).First(req.Context())
	if err != nil {
		log.Fatal(err)
		tst.Error(err)
	}
	if checkInvItem == nil {
		tst.Error("Error: Item doesn't exist")
	} else if checkInvItem.Quantity != 10420690 {
		tst.Error("Error: Item didn't get updated correctly")
		tst.Logf("Item Quantity: %v", checkInvItem.Quantity)
		tst.Logf("idVal: %v", idVal)
		tst.Logf("Item Name: %v, Item ID: %v", checkInvItem.Name, checkInvItem.ID)
	}

}

func TestInvDelete(tst *testing.T) {
	srv := setupServer()

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/store/1391/inventory/?item=%v", idVal), nil)

	rr := executeRequest(req, srv)

	if http.StatusOK != rr.Code {
		tst.Errorf("Error: Expected response code %d. Got %d\n", http.StatusOK, rr.Code)
		tst.Logf("Hint: Check Database if Item_id=%v Entry existed", idVal)
	}

	checkInvItem, _ := srv.DBClient.Item.Query().Where(item.ID(idVal)).First(req.Context())
	if checkInvItem != nil {
		tst.Logf("Item Name: %v, Item ID: %v", checkInvItem.Name, checkInvItem.ID)
		tst.Error("Error: Item not Deleted")
	}
}

func TestGeneratingData(tst *testing.T) {
	/*srv := setupServer()

	// Create the store as you do currently
	store := ent.Store{
		StoreName:    "Gourmet Pantry Delights!",
		OwnerEmail:   "retailgoco@gmail.com",
		StoreAddress: "1234 Main St, Santa Cruz, CA 95060",
		StorePhone:   "831-555-5555",
		StoreType:    "Grocery",
	}

	storeBytes, err := json.Marshal(store)
	if err != nil {
		log.Fatal(err)
		tst.Error(err)
		return
	}

	req, _ := http.NewRequest("POST", "/create/store", bytes.NewBuffer(storeBytes))
	rr := executeRequest(req, srv)
	checkResponseCode(tst, http.StatusCreated, rr.Code)

	// Extract store ID from the response
	var createdStore ent.Store
	if err := json.Unmarshal(rr.Body.Bytes(), &createdStore); err != nil {
		tst.Error("Failed to unmarshal response to get store ID")
		return
	}
	storeID := createdStore.ID // Assuming ID is a field in your Store struct

	// Generate Dummy Inventory Data
	numberOfItems := 4          // Total number of dummy items to create
	categoryChangeInterval := 2 // Change category every 5 items
	var currentCategory string

	for i := 0; i < numberOfItems; i++ {
		// Generate new category name every 5 items
		if i%categoryChangeInterval == 0 {
			currentCategory = gofakeit.BeerStyle() // Using BeerStyle as an example category
		}

		item := ent.Item{
			Name:         gofakeit.BeerName(),
			Photo:        "https://example.com/photo.jpg", // Placeholder
			Quantity:     gofakeit.Number(1, 100),
			Price:        gofakeit.Price(1.00, 1000.00),
			CategoryName: currentCategory,
			StoreID:      storeID,
		}
		itemBytes, err := json.Marshal(item)
		if err != nil {
			tst.Error("Failed to marshal item")
			continue
		}

		itemReq, _ := http.NewRequest("POST", fmt.Sprintf("/store/%d/inventory/create", storeID), bytes.NewBuffer(itemBytes))
		itemRr := executeRequest(itemReq, srv)
		checkResponseCode(tst, http.StatusCreated, itemRr.Code)
	}*/
}

// executeRequest, creates a new ResponseRecorder
// then executes the request by calling ServeHTTP in the router
// after which the handler writes the response to the response recorder
// which we can then inspect.
func executeRequest(req *http.Request, s *server.Server) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode is a simple utility to check the response code
// of the response
func checkResponseCode(tst *testing.T, expected, actual int) {
	if expected != actual {
		tst.Errorf("Error: Expected response code %d. Got %d\n", expected, actual)
		tst.Log("Hint: Check Database if Item Entry existed")
	}
}
func setupServer() *server.Server {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.Server{
		Router:   chi.NewRouter(),
		DBClient: util.Open(&config),
		Config:   &config,
	}

	srv.MountHandlers()

	return &srv
}
