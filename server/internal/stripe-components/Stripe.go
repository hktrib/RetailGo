package StripeComponents

import (
	"encoding/json"
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/account"
	"github.com/stripe/stripe-go/v76/accountlink"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/price"
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

// CreateStripeItem creates a new product in Stripe
func CreateStripeItem(item *ent.Item, store *ent.Store) (*stripe.Product, error) {

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

func UpdateStripeItem(item *ent.Item, name string, StoreStripeID string) (*stripe.Product, error) {

	productParams := &stripe.ProductParams{}
	productParams.Name = stripe.String(name)
	product, err := product.Update(item.StripeProductID, productParams)

	if err != nil {
		return nil, err
	}
	return product, nil
}
func UpdateStripePrice(item *ent.Item, newPrice float64, StoreStripeID string) (*stripe.Price, error) {

	priceParams := &stripe.PriceParams{
		Product:    stripe.String(item.StripeProductID),
		Currency:   stripe.String(string(stripe.CurrencyUSD)),
		UnitAmount: stripe.Int64(int64(newPrice * 100)),
	}
	priceId, err := price.New(priceParams)
	if err != nil {
		return nil, err
	}

	productParams := &stripe.ProductParams{DefaultPrice: stripe.String(priceId.ID)}
	_, err = product.Update(item.StripeProductID, productParams)
	if err != nil {
		return nil, err
	}
	return priceId, nil
}

func CreateCheckoutSession(items []CartItem, StoreStripeID string, StoreId int, w http.ResponseWriter, r *http.Request) {
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
		PaymentIntentData: &stripe.CheckoutSessionPaymentIntentDataParams{
			TransferData: &stripe.CheckoutSessionPaymentIntentDataTransferDataParams{
				Destination: stripe.String(StoreStripeID),
			},
			ApplicationFeeAmount: stripe.Int64(500),
		},

		LineItems: lineItems,
		Mode:      stripe.String(string(stripe.CheckoutSessionModePayment)),
		ReturnURL: stripe.String(fmt.Sprintf("https://retail-go.vercel.app/store/%d/pos", StoreId)),
	}

	s, err := session.New(params)
	fmt.Printf("session.New: %v", s)

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
func StartOnboarding(accountId string, storeID int) (*stripe.AccountLink, error) {

	// Set your secret key. Remember to switch to your live secret key in production.
	// See your keys here: https://dashboard.stripe.com/apikeys

	params := &stripe.AccountLinkParams{
		Account:    stripe.String(accountId),
		RefreshURL: stripe.String(fmt.Sprintf("https://retailgo-production.up.railway.app/store/%d/onboarding", storeID)),
		ReturnURL:  stripe.String("https://retail-go.vercel.app/store/"),
		Type:       stripe.String("account_onboarding"),
		Collect:    stripe.String("eventually_due"),
	}
	result, err := accountlink.New(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
