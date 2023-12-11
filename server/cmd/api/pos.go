package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"

	StripeHelper "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
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
	// get store id from url query string
	store_id, id_err := strconv.Atoi(chi.URLParam(request, "store_id"))
	if id_err != nil {
		log.Debug().Err(err).Msg("StoreCheckout: unable to parse store_id")
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// get store
	targetStore, err := srv.DBClient.Store.Query().Where(store.ID(store_id)).Only(request.Context())
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
	StripeHelper.CreateCheckoutSession(cart, targetStore.StripeAccountID, writer, request)

}

/*
StripeWebhookRouter Brief:

-> Handles incoming webhook events from Stripe and routes them based on event types.

	Verifies the webhook signature, constructs the event, and routes it to respective
	functions based on event types such as account updates, payment successes/failures, etc.

External Package Calls:
- webhook.ConstructEvent()
- HandleTransSuccess(w, event, srv) [for specific event types]
*/
func (srv *Server) StripeWebhookRouter(w http.ResponseWriter, r *http.Request, endpointSecret string) {

	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Debug().Err(err).Msg("StripeWebhookRouter: server failed to read request body")
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// This is your Stripe CLI webhook secret for testing your endpoint locally.

	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, r.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "account.updated":
		// Then define and call a function to handle the event account.updated
	case "account.application.authorized":
		// Then define and call a function to handle the event account.application.authorized
	case "account.application.deauthorized":
		// Then define and call a function to handle the event account.application.deauthorized
	case "account.external_account.created":
		// Then define and call a function to handle the event account.external_account.created
	case "account.external_account.deleted":
		// Then define and call a function to handle the event account.external_account.deleted
	case "account.external_account.updated":
		// Then define and call a function to handle the event account.external_account.updated
	case "checkout.session.async_payment_failed":
		// Then define and call a function to handle the event checkout.session.async_payment_failed
	case "checkout.session.async_payment_succeeded":
		// Then define and call a function to handle the event checkout.session.async_payment_succeeded
		HandleTransSuccess(w, event, srv)

	case "checkout.session.completed":
		HandleTransSuccess(w, event, srv)
		// Then define and call a function to handle the event checkout.session.expired
	case "product.created":
		// Then define and call a function to handle the event product.created
	case "product.deleted":
		// Then define and call a function to handle the event product.deleted
	case "product.updated":
		// Then define and call a function to handle the event product.updated
	// ... handle other event types
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)

}

/*
HandleTransSuccess Brief:

-> Handles successful transaction events from Stripe webhook.

	Parses the webhook JSON, retrieves line items from the session, and fulfills orders by updating
	item quantities using srv.FulfillOrder(LineItemList).

External Package Calls:
- json.Unmarshal()
- srv.FulfillOrder()
*/
func HandleTransSuccess(w http.ResponseWriter, event stripe.Event, srv *Server) bool {

	var session stripe.CheckoutSession
	err := json.Unmarshal(event.Data.Raw, &session)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}

	params := &stripe.CheckoutSessionParams{}
	params.AddExpand("line_items")

	// Retrieve the session. If you require line items in the response, you may include them by expanding line_items.
	lineItems := session.LineItems
	// Fulfill the purchase...
	srv.FulfillOrder(lineItems)

	return false
}

/*
FulfillOrder Brief:

-> Updates item quantities based on the provided LineItemList from Stripe.

External Package Calls:
- srv.DBClient.Item.Query()
- srv.DBClient.Item.UpdateOne()
*/
func (srv *Server) FulfillOrder(LineItemList *stripe.LineItemList) {
	for i := range LineItemList.Data {
		// update item quantity

		LineItem, err := srv.DBClient.Item.Query().Where(item.StripeProductID(LineItemList.Data[i].ID)).Only(context.Background())
		if err != nil {
			panic(err)
		}
		_, err = srv.DBClient.Item.
			UpdateOne(LineItem).
			SetQuantity(LineItem.Quantity - int(LineItemList.Data[i].Quantity)).
			AddNumberSoldSinceUpdate(int(LineItemList.Data[i].Quantity)).
			SetDateLastSold(time.Now().Format("2006-01-02")).
			Save(context.Background())

		if err != nil {
			log.Debug().Err(err).Msg("FulfillOrder: Unable to update item")
		}
	}
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
	response := make(map[string][]interface{}, 1)
	response["categories"] = make([]interface{}, 0)
	response["items"] = make([]interface{}, 0)
	for i, cat := range categories {
		// get items for each category
		c, _ := json.Marshal(cat) // marshal the category into a map
		catInfo := make(map[string]interface{})
		_ = json.Unmarshal(c, &catInfo) // unmarshal the map into the response map
		response["categories"] = append(response["categories"], catInfo)
		delete(response["categories"][i].(map[string]interface{}), "edges") // the edges field is not needed in the response
		items, _ := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(cat.ID))).All(ctx)
		pitems := PruneItems(items...)
		for _, item1 := range pitems {
			response["items"] = append(response["items"], item1)
		}
	}
	responseBody, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
