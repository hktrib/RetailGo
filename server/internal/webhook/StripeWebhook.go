package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	StripeHelper "github.com/hktrib/RetailGo/cmd/api/stripe-components"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/webhook"
	"io"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"

	"bytes"

	"html/template"
	"os"
	"time"
)

type HtmlTemplate struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Store_name  string `json:"store_name"`
	Sender_name string `json:"sender_name"`
	Action_url  string `json:"action_url"`
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
		AuthorizeUser(w, event, DBClient)
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
	default:
		fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		w.WriteHeader(http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func AuthorizeUser(w http.ResponseWriter, event stripe.Event, client *ent.Client) {
	var CAccount stripe.Account
	err := json.Unmarshal(event.Data.Raw, &CAccount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the store from the database
	Store, err := client.Store.Query().Where(store.StripeAccountID(CAccount.ID)).Only(context.Background())

	// Update the user's status to authorized
	_, err = client.Store.UpdateOne(Store).SetIsAuthorized(true).Save(context.Background())

	if err != nil {
		log.Debug().Err(err).Msg("AuthorizeUser: unable to update store")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = SendSuccessEmail(Store)
	if err != nil {
		return
	}

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
func SendSuccessEmail(StoreObj *ent.Store) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"retailgoco@gmail.com",
		"hhqm pqgw lxmi weqb",
		"smtp.gmail.com",
	)

	// Define email headers and HTML content
	from := mail.Address{Name: "RetailGo Team", Address: "retailgoco@gmail.com"}
	em := StoreObj.OwnerEmail
	split := strings.Split(em, "@")

	to := mail.Address{Name: split[0], Address: StoreObj.OwnerEmail}
	subject := "You must complete onboarding to continue!"

	// Message to be sent
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	// HTML message

	tmpl := template.Must(template.ParseGlob("./cmd/templates/*"))

	url, err := StripeHelper.StartOnboarding(StoreObj.StripeAccountID)
	if err != nil {
		return fmt.Errorf("failed to start onboarding: %w", err)
	}

	htmlBody := new(bytes.Buffer)
	templateData := HtmlTemplate{
		Action_url: url.URL,
	}

	err = tmpl.ExecuteTemplate(htmlBody, "onboarding_success.html", templateData)
	if err != nil {
		return fmt.Errorf("failed to read email template: %w", err)
	}

	// Combine headers and HTML content
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody.String()

	// Convert to and from to slice of strings
	toAddresses := []string{to.Address}

	// Send the email
	err = smtp.SendMail("smtp.gmail.com:587", auth, from.Address, toAddresses, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
