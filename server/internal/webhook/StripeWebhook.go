package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
	"io"
	"net/http"
	"os"
	"time"
)

/*
StripeWebhookRouter Brief:

-> Handles incoming webhook events from Stripe and routes them based on event types.

	Verifies the webhook signature, constructs the event, and routes it to respective
	functions based on event types such as account updates, payment successes/failures, etc.

External Package Calls:
- webhook.ConstructEvent()
- HandleTransSuccess(w, event, srv) [for specific event types]
*/
func StripeWebhookRouter(w http.ResponseWriter, r *http.Request, endpointSecret string, DBClient *ent.Client) {

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
		HandleTransSuccess(w, event, DBClient)

	case "checkout.session.completed":
		HandleTransSuccess(w, event, DBClient)
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
func HandleTransSuccess(w http.ResponseWriter, event stripe.Event, DBClient *ent.Client) bool {

	var CSession stripe.CheckoutSession
	err := json.Unmarshal(event.Data.Raw, &CSession)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return true
	}

	params := &stripe.CheckoutSessionParams{}
	params.AddExpand("line_items")

	// Retrieve the session.
	sessionWithLineItems, err := session.Get(CSession.ID, params)
	if err != nil {
		log.Debug().Err(err).Msg("HandleTransSuccess: unable to retrieve session")
		w.WriteHeader(http.StatusBadRequest)
		return true
	}
	lineItems := sessionWithLineItems.LineItems
	// Fulfill the purchase...
	FulfillOrder(lineItems, DBClient)

	return false
}

/*
FulfillOrder Brief:

-> Updates item quantities based on the provided LineItemList from Stripe.

External Package Calls:
- srv.DBClient.Item.Query()
- srv.DBClient.Item.UpdateOne()
*/
func FulfillOrder(LineItemList *stripe.LineItemList, DBClient *ent.Client) {
	for i := range LineItemList.Data {
		// update item quantity

		LineItem, err := DBClient.Item.Query().Where(item.StripePriceID(LineItemList.Data[i].Price.ID)).Only(context.Background())
		if err != nil {
			log.Debug().Err(err).Msg("FulfillOrder: Unable to retrieve item")
			fmt.Printf("%v+", LineItemList.Data[i].ID)
		}
		_, err = DBClient.Item.
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
