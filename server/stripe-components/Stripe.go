package StripeComponents

import (
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/product"
	"net/http"
)

type CartItem struct {
	Id       int
	Item     *ent.Item
	Quantity int
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
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems:  lineItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String("http://localhost:8080/" + "success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("http://localhost:8080/" + "cancel?session_id={CHECKOUT_SESSION_ID}"),
	}

	s, err := session.New(params)

	if err != nil {
		fmt.Printf("session.New: %v", err)
	}
	w.WriteHeader(302)
	fmt.Fprintf(w, "%s", s.URL)

}
