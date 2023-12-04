package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	clerkHelpers "github.com/hktrib/RetailGo/internal/clerk"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
	"github.com/hktrib/RetailGo/internal/ent/usertostore"
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
		SetClerkUserID(reqUser.ClerkUserID).
		SetEmail(reqUser.Email).
		SetIsOwner(reqUser.IsOwner).
		SetFirstName(reqUser.FirstName).
		SetLastName(reqUser.LastName).
		SetStoreID(reqUser.StoreID).Save(ctx)
	
	if err != nil {
		log.Debug().Err(err).Msg("User Creation didn't work:")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Add StoreID -> to "stores" list in private_metadata for clerk user

	clerkStore, err := clerkHelpers.NewClerkStore(srv.ClerkClient, reqUser.ClerkUserID, srv.Config)
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

	createdItem, create_err := srv.DBClient.User.Create().SetFirstName(req.FirstName).SetLastName(req.LastName).SetEmail(req.Email).SetIsOwner(false).SetStoreID(req.StoreID).Save(ctx)

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
	if name, ok := values["first_name"]; ok {
		targetUser.Update().SetFirstName(name[0]).SaveX(ctx)
	}
	if name, ok := values["last_name"]; ok {
		targetUser.Update().SetLastName(name[0]).SaveX(ctx)
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

func (srv *Server) UserHasStore(w http.ResponseWriter, r *http.Request) {

	// Verify user credentials using clerk
	clerk_user, err := VerifyUserCredentials(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

		return
	}
	ctx := r.Context()

	// Extract the Clerk user ID
	clerkID := clerk_user.ID

	// Query the User table for a match with the Clerk ID
	matchedUser, err := srv.DBClient.User.Query().Where(user.ClerkUserID(clerkID)).Only(ctx)
	if ent.IsNotFound(err) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Query the UserToStore table using the found user's ID
	_, err = srv.DBClient.UserToStore.Query().Where(usertostore.UserID(matchedUser.ID)).Only(ctx)

	if ent.IsNotFound(err) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("false"))
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// If no error, it means a record was found
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("true"))
}
func (srv *Server) UserJoinStore(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Verify user credentials using clerk
	clerkUser, clerkErr := VerifyUserCredentials(r)
	if clerkErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var reqData struct {
		ClerkID string `json:"ClerkUserID"`
		StoreId string `json:"StoreId"`
	}
	err = json.Unmarshal(reqBody, &reqData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if reqData.ClerkID != clerkUser.ID {
		http.Error(w, "You can only accept the invite on your behalf.\nclerkId = "+clerkUser.ID+"\n reqId: "+reqData.ClerkID, http.StatusBadRequest)
		return
	}

	var firstName, lastName, email string
	if clerkUser.FirstName != nil {
		firstName = *clerkUser.FirstName
	} else {
		// Handle the nil case for FirstName
	}

	if clerkUser.LastName != nil {
		lastName = *clerkUser.LastName
	} else {
		// Handle the nil case for LastName
	}

	if clerkUser.EmailAddresses != nil {
		email = clerkUser.EmailAddresses[0].EmailAddress
	}

	store, err := srv.DBClient.Store.Query().Where(store.UUID(reqData.StoreId)).Only(ctx)
	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	} // Create a new User entity

	_, create_err := srv.DBClient.User.Create().
		SetEmail(email).
		SetIsOwner(false).
		SetStoreID(store.ID).
		SetClerkUserID(clerkUser.ID).
		SetFirstName(firstName).
		SetLastName(lastName).
		Save(ctx)

	if create_err != nil {
		fmt.Println("User Creation didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully\n"))
}
