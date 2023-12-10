package StripeComponents

import (
	"encoding/json"
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/product"
	"net/http"
)

type CartItem struct {
	Id       int
	Item     *ent.Item
	Quantity int
}

func CreateConnectedAccount() (*stripe.Account, error) {

	params := &stripe.AccountParams{
		Type: stripe.String(string(stripe.AccountTypeExpress)),
		Capabilities: &stripe.AccountCapabilitiesParams{
			CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
				Requested: stripe.Bool(true),
			},
			Transfers: &stripe.AccountCapabilitiesTransfersParams{
				Requested: stripe.Bool(true),
			},
		},
	}
	result, err := account.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func startOnboarding(accountId string) (*stripe.AccountLink, error) {

	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys

	params := &stripe.AccountLinkParams{
		Account:    stripe.String("{{CONNECTED_ACCOUNT_ID}}"),
		RefreshURL: stripe.String("https://example.com/reauth"),
		ReturnURL:  stripe.String("https://example.com/return"),
		Type:       stripe.String("account_onboarding"),
		Collect:    stripe.String("eventually_due"),
	}
	result, err := accountlink.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateStripeItem creates a new product in Stripe
func CreateStripeItem(item *ent.Item) (*stripe.Product, error) {

	productParams := &stripe.ProductParams{
		Name: stripe.String(item.Name),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			UnitAmount: stripe.Int64(int64(item.Price * 100)),
		},
	}
	product, err := product.New(productParams)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func UpdateStripeItem(item *ent.Item, price float64, name string) (*stripe.Product, error) {

	productParams := &stripe.ProductParams{
		Name: stripe.String(name),
		DefaultPriceData: &stripe.ProductDefaultPriceDataParams{
			Currency:   stripe.String(string(stripe.CurrencyUSD)),
			UnitAmount: stripe.Int64(int64(price * 100)),
		},
	}
	product, err := product.Update(item.StripeProductID, productParams)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func CreateCheckoutSession(items []CartItem, w http.ResponseWriter, r *http.Request) {
	//TODO: ADD ENVIRONMENT VARIABLE FOR SERVER ADDRESS
	var lineItems []*stripe.CheckoutSessionLineItemParams

	for i := range items {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(items[i].Item.StripePriceID),
			Quantity: stripe.Int64(int64(items[i].Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		UIMode: stripe.String("embedded"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems: lineItems,
		Mode:      stripe.String(string(stripe.CheckoutSessionModePayment)),
		ReturnURL: stripe.String("https://retail-go.vercel.app/store"),
	}

	s, err := session.New(params)

	if err != nil {
		fmt.Printf("session.New: %v", err)
	}
	w.WriteHeader(200)
	res := map[string]string{
		"ClientSecret": s.ClientSecret,
	}
	resp, _ := json.Marshal(res)
	w.Write(resp)
}

func HandleOnboarding(w http.ResponseWriter, r *http.Request) {
	store_id := r.Context().Value("store_var").(int)
	// Get store
	store, err := ent.FromContext(r.Context()).Store.Query().Where(store.IDEQ(store_id)).Only(r.Context())
	if err != nil {
		log.Debug().Err(err).Msg("HandleOnboarding: unable to fetch store from database")
		return
	}
	accLink, err := startOnboarding(store.StripeAccountID)
	if err != nil {
		log.Debug().Err(err).Msg("HandleOnboarding: unable to start onboarding")
		return
	}
	w.WriteHeader(302)
	fmt.Fprintf(w, "%s", accLink.URL)

}
