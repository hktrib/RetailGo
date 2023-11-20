package server

import (
	"encoding/json"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/item"
	. "github.com/hktrib/RetailGo/stripe-components"
	"io"
	"net/http"
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
