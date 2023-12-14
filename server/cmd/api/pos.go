package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	StripeHelper "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strconv"
)

/*
StoreCheckout Brief:

-> Processes a checkout for items in a cart.

	Reads request body to retrieve cart items, queries the database for each item,
	and creates a Stripe Checkout Session using S

External Package Calls:
- srv.DBClient.Item.Query()
- StripeHelper.CreateCheckoutSession()
*/
func (srv *Server) StoreCheckout(writer http.ResponseWriter, request *http.Request) {
	var cart []StripeHelper.CartItem
	req_body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Debug().Err(err).Msg("StoreCheckout: unable to read request body")
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(req_body, &cart)
	if err != nil {
		log.Debug().Err(err).Msg("StoreCheckout: unable to Unmarshal request body")
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// query for items
	for i := range cart {
		LineItem, err := srv.DBClient.Item.Query().Where(item.ID(cart[i].Id)).Only(request.Context())
		if ent.IsNotFound(err) {
			log.Debug().Err(err).Msg("StoreCheckout: ent didn't find the item in the database")
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if err != nil {
			log.Debug().Err(err).Msg("StoreCheckout: server failed to execute find item query")
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		cart[i].Item = LineItem

	}
	// get store id from url
	storeId, err := strconv.Atoi(chi.URLParam(request, "store_id"))
	if err != nil {
		log.Debug().Err(err).Msg("StoreCheckout: unable to parse storeId")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// get store
	targetStore, err := srv.DBClient.Store.Query().Where(store.ID(storeId)).Only(request.Context())
	if ent.IsNotFound(err) {
		log.Debug().Err(err).Msg("StoreCheckout: ent didn't find the store in the database")
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Debug().Err(err).Msg("StoreCheckout: server failed to execute find store query")
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create a new Stripe Checkout Session
	StripeHelper.CreateCheckoutSession(cart, targetStore.StripeAccountID, targetStore.ID, writer, request)

}

/*
GetPosInfo Brief:

-> Retrieves Information to be displayed on the stores Pos interface.

External Package Calls:
- srv.DBClient.Category.Query()
- srv.DBClient.Item.Query()
*/
func (srv *Server) GetPosInfo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	categories, err := srv.DBClient.Category.Query().Where(category.StoreID(store_id)).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("GetPosInfo failed")
	}
	response := make(map[string][]interface{}, 2)
	response["categories"] = make([]interface{}, 0)
	response["items"] = make([]interface{}, 5)
	for i, cat := range categories {
		// get items for each category
		c, _ := json.Marshal(cat) // marshal the category into a map
		catInfo := make(map[string]interface{})
		_ = json.Unmarshal(c, &catInfo) // unmarshal the map into the response map
		response["categories"] = append(response["categories"], catInfo)
		delete(response["categories"][i].(map[string]interface{}), "edges") // the edges field is not needed in the response
		items, _ := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(cat.ID))).All(ctx)
		prunedItems := PruneItems(items...)
		for _, item1 := range prunedItems {
			response["items"] = append(response["items"], item1)
		}
	}
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
