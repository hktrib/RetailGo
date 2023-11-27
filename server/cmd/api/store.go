package server

import (
	"fmt"
	"net/http"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/transactions"
)

func (srv *Server) CreateStore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	reqStore := ctx.Value(Param("store_var")).(*ent.Store)
	reqUser := ctx.Value(Param("owner")).(*ent.User)
	err := transactions.StoreAndOwnerCreationTx(ctx, reqStore, reqUser, srv.DBClient)
	if err != nil {
		fmt.Println("Could not executed Store|Owner Transaction:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created!"))
}
