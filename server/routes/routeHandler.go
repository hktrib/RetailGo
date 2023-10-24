package routes

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent"
	store2 "github.com/hktrib/RetailGo/ent/store"
	"github.com/hktrib/RetailGo/util"
	"net/http"
	"strconv"
)

type RouteHandler struct {
	Client *ent.Client
	Config *util.Config
}

func (rh *RouteHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World -> RetailGo!!"))

	ctx := r.Context()

	user, err := rh.Client.User.
		Create().
		SetUsername("apsingh").
		SetRealName("Apramay Singh").
		SetEmail("apramaysingh@gmail.com").
		SetIsOwner(true).
		SetStoreID(1321).
		Save(ctx)

	store, err := rh.Client.Store.
		Create().
		SetStoreName("Retail Go").SetID(1321).
		Save(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(store)
}
func (rh *RouteHandler) ValidateStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store_id := chi.URLParam(r, "store_id")
		StoreId, err := strconv.Atoi(store_id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		store, err := rh.Client.Store.
			Query().Where(store2.ID(StoreId)).Only(r.Context())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusInternalServerError)
			return
		}
		if store == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "store_var", r.Context())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
