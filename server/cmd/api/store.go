package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	clerkHelpers "github.com/hktrib/RetailGo/internal/clerk"
	"github.com/hktrib/RetailGo/internal/ent/user"
	"github.com/rs/zerolog/log"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/transactions"
)

func (srv *Server) CreateStore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	reqStore := ctx.Value(Param("store")).(*ent.Store)
	reqUser := ctx.Value(Param("owner")).(*ent.User)
	err := transactions.StoreAndOwnerCreationTx(ctx, reqStore, reqUser, srv.DBClient)
	if err != nil {
		fmt.Println("Could not executed Store|Owner Transaction:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Adding storeid to clerk store
	clerkStore, err := clerkHelpers.NewClerkStore(srv.ClerkClient, reqUser.ClerkUserID)
	if err != nil {
		log.Debug().Err(err).Msg("NewClerkStore failed: Unable to create ClerkStore instance using clerk user id:")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return	
	}

	err = clerkStore.AddStore(reqUser.StoreID)
	if err != nil {
		log.Debug().Err(err).Msg("AddStore failed: ")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return	
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created!"))
}
func (srv *Server) GetStoreUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	StoreID, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	Users := srv.DBClient.User.Query().Where(user.StoreID(StoreID)).AllX(ctx)
	resp, err := json.Marshal(Users)

	var userInfo []map[string]interface{}
	err = json.Unmarshal(resp, &userInfo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for i := range userInfo {
		delete(userInfo[i], "clerk_user_id")
		delete(userInfo[i], "store_id")
		delete(userInfo[i], "edges")

	}

	resp, err = json.Marshal(userInfo)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
