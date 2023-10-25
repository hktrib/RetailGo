package routes

import (
	"fmt"
	"net/http"

	"github.com/hktrib/RetailGo/ent"
	"github.com/hktrib/RetailGo/util"
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

	if err != nil {
		fmt.Println(err)
	}

	store, err := rh.Client.Store.
		Create().
		SetStoreName("Retail Go").
		Save(ctx)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	fmt.Println(store)
}
