package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/user"
)

func (srv *Server) UserCreate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	reqUser := ent.User{}

	err = json.Unmarshal(reqBody, &reqUser)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = srv.DBClient.User.Create().
		SetEmail(reqUser.Email).
		SetIsOwner(reqUser.IsOwner).
		SetRealName(reqUser.RealName).
		SetStoreID(reqUser.StoreID).Save(ctx)

	if err != nil {
		fmt.Println("User Creation didn't work:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK\n"))
}

func (srv *Server) UserDelete(w http.ResponseWriter, r *http.Request) {

	user_id, err := strconv.Atoi(r.URL.Query().Get("user"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = srv.DBClient.
		User.DeleteOneID(user_id).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (srv *Server) UserQuery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user_data, err := srv.DBClient.User.Query().Where(user.ID(user_id)).Only(ctx)
	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res, err := json.MarshalIndent(user_data, "", "")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func (srv *Server) userAdd(w http.ResponseWriter, r *http.Request) {
	// name, photo, quantity, store_id, category

	ctx := r.Context()

	// Load store_id first

	// Load the message parameters (Name, Photo, Quantity, Category)
	req_body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Server failed to read message body:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	req := ent.User{}

	body_parse_err := json.Unmarshal(req_body, &req)

	if body_parse_err != nil {
		fmt.Println("Unmarshalling failed:", body_parse_err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	createdItem, create_err := srv.DBClient.User.Create().SetUsername(req.Username).SetRealName(req.RealName).SetEmail(req.Email).SetIsOwner(false).SetStoreID(req.StoreID).Save(ctx)

	if create_err != nil {
		fmt.Println("Create didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"id": createdItem.ID,
	})

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}

func (srv *Server) userUpdate(w http.ResponseWriter, r *http.Request) {
	// get item id from url query string
	ctx := r.Context()
	user_id, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetUser, err := srv.DBClient.User.Query().Where(user.ID(user_id)).Only(ctx)

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//get all query params
	values := r.URL.Query()

	// set appropriate fields
	if name, ok := values["username"]; ok {
		targetUser.Update().SetUsername(name[0]).SaveX(ctx)
	}
	if name, ok := values["realname"]; ok {
		targetUser.Update().SetRealName(name[0]).SaveX(ctx)
	}
	if name, ok := values["email"]; ok {
		targetUser.Update().SetEmail(name[0]).SaveX(ctx)
	}
	if name, ok := values["isowner"]; ok {
		targetUser.Update().SetIsOwner(name[0] == "true").SaveX(ctx)
	}
	if name, ok := values["storeid"]; ok {
		val, err := strconv.Atoi(name[0])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		targetUser.Update().SetStoreID(val).SaveX(ctx)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("OK"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
