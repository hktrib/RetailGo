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
		Photo:    []byte("uhsrgouhsrgouhsrgoushrg"),
		Quantity: 15,
		Price:    120.9,
		StoreID:  1391,
	}

	invItemBytes, err := json.Marshal(invItem)
	bytes := bytes.NewBuffer(invItemBytes)
	if err != nil {
		log.Fatal(err)
		tst.Error(err)
	}

	req, _ := http.NewRequest("POST", "/store/1391/inventory/create", bytes)

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
	config, err := util.LoadConfig("../")
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
