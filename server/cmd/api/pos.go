package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func (srv *Server) StoreCheckout(writer http.ResponseWriter, request *http.Request) {
	var cart []CartItem
	req_body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(req_body, &cart)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// query for items
	for i := range cart {
		LineItem, err := srv.DBClient.Item.Query().Where(item.ID(cart[i].Id)).Only(request.Context())
		if ent.IsNotFound(err) {
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		cart[i].Item = LineItem

	}

	// Create a new Stripe Checkout Session
	CreateCheckoutSession(cart, writer, request)

}
func (srv *Server) StripeWebhookRouter(w http.ResponseWriter, r *http.Request) {

	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	endpointSecret := srv.Config.STRIPE_WEBHOOK_SECRET
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

func (srv *Server) FulfillOrder(LineItemList *stripe.LineItemList) {
	for i := range LineItemList.Data {
		// update item quantity
		LineItem, err := srv.DBClient.Item.Query().Where(item.StripeProductID(LineItemList.Data[i].ID)).Only(context.Background())
		if err != nil {
			panic(err)
		}
		_, err = srv.DBClient.Item.UpdateOne(LineItem).SetQuantity(LineItem.Quantity - int(LineItemList.Data[i].Quantity)).Save(context.Background())

	}
}

func (srv *Server) GetPosInfo(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		fmt.Println("Invalid store id:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	categories, err := srv.DBClient.Category.Query().Where(category.StoreID(store_id)).All(ctx)
	response := make(map[string][]interface{}, 1)

	for i, cat := range categories {
		// get items for each category
		c, _ := json.Marshal(cat) // marshal the category into a map
		catInfo := make(map[string]interface{})
		_ = json.Unmarshal(c, &catInfo) // unmarshal the map into the response map
		response["categories"] = append(response["categories"], catInfo)
		delete(response["categories"][i].(map[string]interface{}), "edges") // the edges field is not needed in the response
		items, _ := srv.DBClient.Item.Query().Where(item.HasCategoryWith(category.ID(cat.ID))).All(ctx)
		for j, item2 := range items {
			it, _ := json.Marshal(item2)
			itemInfo := make(map[string]interface{})
			_ = json.Unmarshal(it, &itemInfo) // unmarshal the map into the respons
			response["items"] = append(response["items"], itemInfo)
			delete(response["items"][j].(map[string]interface{}), "edges")
			response["items"][j].(map[string]interface{})["category_id"] = cat.ID // these typecasts are necessary go  doesn't allow you to add or remove fields to the struct
			response["items"][j].(map[string]interface{})["category"] = cat.Name

		}
		responseBody, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}
}
