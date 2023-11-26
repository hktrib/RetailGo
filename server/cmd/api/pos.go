package server

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/webhook"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
func (srv *Server) HandleSuccess(w http.ResponseWriter, r *http.Request) {

	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// Pass the request body and Stripe-Signature header to ConstructEvent, along with the webhook signing key
	// You can find your endpoint's secret in your webhook settings
	endpointSecret := "whsec_..."
	event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}
	fmt.Println("Webhook received!")
	fmt.Printf("%+v\n", event)
	// Handle the checkout.session.completed event
	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &session)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		params := &stripe.CheckoutSessionParams{}
		params.AddExpand("line_items")

		// Retrieve the session. If you require line items in the response, you may include them by expanding line_items.
		lineItems := session.LineItems
		// Fulfill the purchase...
		srv.FulfillOrder(lineItems)
	}

	w.WriteHeader(http.StatusOK)

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
