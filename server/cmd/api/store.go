package server

import (
	"fmt"
	"net/http"

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
