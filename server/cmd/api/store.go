package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/internal/ent/user"

	"github.com/hktrib/RetailGo/internal/transactions"
)

func (srv *Server) CreateStore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	err := transactions.StoreAndOwnerCreationTx(ctx, srv.DBClient)
	if err != nil {
		fmt.Println("Could not executed Store|Owner Transaction:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// reqBody, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("Server failed to read message body:", err)
	// 	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 	return
	// }

	// reqItem := ent.Store{}

	// err = json.Unmarshal(reqBody, &reqItem)
	// if err != nil {
	// 	fmt.Println("Unmarshalling failed:", err)
	// 	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	// 	return
	// }
	// createdItem, create_err := srv.DBClient.Store.Create().
	// 	SetStoreName(reqItem.StoreName).Save(ctx)

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
	}

	resp, err = json.Marshal(userInfo)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
