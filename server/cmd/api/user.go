package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	clerkstorage "github.com/hktrib/RetailGo/internal/clerk"
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

	newUser, err := srv.DBClient.User.Create().
		SetClerkUserID(reqUser.ClerkUserID).
		SetEmail(reqUser.Email).
		SetIsOwner(reqUser.IsOwner).
		SetFirstName(reqUser.FirstName).
		SetLastName(reqUser.LastName).
		SetStoreID(reqUser.StoreID).Save(ctx)

	if err != nil {
		log.Debug().Err(err).Msg("User Creation didn't work")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Creating New Clerk Store instance of the given User
	clerkStore, err := clerkstorage.NewClerkStore(srv.ClerkClient, reqUser.ClerkUserID, srv.Config)
	if err != nil {
		log.Debug().Err(err).Msg("NewClerkStore failed: Unable to create ClerkStore instance using clerk user id")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Fetching the UserToStore Relation from the database
	newUTS, err := srv.DBClient.UserToStore.Query().
		Where(usertostore.StoreID(newUser.StoreID), usertostore.UserID(newUser.ID)).
		Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Store Query for Clerk Store updating didn't work")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Adding the
	err = clerkStore.AddMetadata(newUTS)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to add user metadata to Clerk Store")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK\n"))
}

func (srv *Server) UserDelete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Debug().Err(err).Msg("Bad/No user_id")
		http.Error(w, http.StatusText(http.StatusBadRequest)+": no user id", http.StatusBadRequest)
		return
	}

	// Fetching User from Database
	fetchedUser, err := srv.DBClient.User.Query().
		Where(user.ID(userID)).Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to fetch user from database")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Querying for the user to store relation
	newUTS, err := srv.DBClient.UserToStore.Query().
		Where(usertostore.UserID(fetchedUser.ID), usertostore.StoreID(fetchedUser.StoreID)).
		Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to query entry in userToStore table")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete entries in usertostore table based on userID
	_, err = srv.DBClient.
		UserToStore.
		Delete().
		Where(usertostore.UserID(userID)). // Assuming there is a UserID field in userToStore
		Exec(r.Context())
	if err != nil {
		log.Debug().Err(err).Msg("Unable to delete entry in userToStore table")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Deleting user from Database
	err = srv.DBClient.
		User.DeleteOneID(userID).
		Exec(r.Context())

	if ent.IsNotFound(err) {
		log.Debug().Err(err).Msg("ent is not found")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+":"+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create Clerk Store Instance
	clerkStore, err := clerkstorage.NewClerkStore(srv.ClerkClient, fetchedUser.ClerkUserID, srv.Config)
	if err != nil {
		log.Debug().Err(err).Msg("unable to create clerk store instance")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Removing user's store relation from user's Clerk Public Metadata.
	err = clerkStore.RemoveMetadata(newUTS)
	if err != nil {
		log.Debug().Err(err).Msg("unable to remove user's store relation from Clerk Store")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (srv *Server) UserQuery(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user_data, err := srv.DBClient.User.Query().Where(user.ID(userID)).Only(ctx)
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

// Assuming you have a struct to map the request body
type UserUpdateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	// Add other fields as necessary
}

func (srv *Server) userUpdate(w http.ResponseWriter, r *http.Request) {
	// get item id from url query string
	ctx := r.Context()
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	targetUser, err := srv.DBClient.User.Query().Where(user.ID(userID)).Only(ctx)

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound)+": user not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Decode the request body
	var req UserUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Update user with new values
	update := targetUser.Update()
	if req.FirstName != "" {
		update.SetFirstName(req.FirstName)
	}
	if req.LastName != "" {
		update.SetLastName(req.LastName)
	}
	if req.Email != "" {
		update.SetEmail(req.Email)
	}

	// Save the changes
	if _, err := update.Save(ctx); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

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
	}

	// Create a new User entity
	newUser, create_err := srv.DBClient.User.Create().
		SetEmail(email).
		SetIsOwner(false).
		SetStoreID(store.ID).
		SetClerkUserID(clerkUser.ID).
		SetFirstName(firstName).
		SetLastName(lastName).
		AddStoreIDs(store.ID).
		Save(ctx)
	if create_err != nil {
		fmt.Println("User Creation didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, user_store_error := srv.DBClient.UserToStore.Update().
		SetStoreID(store.ID).
		SetStoreName(store.StoreName).
		SetClerkUserID(clerkUser.ID).
		SetPermissionLevel(1). // Employee permission level
		Save(ctx)
	if user_store_error != nil {
		fmt.Println("User Creation didn't work:", create_err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Creating ClerkStore instance for new employee to add user_to_store relationship to metadata
	clerkStore, err := clerkstorage.NewClerkStore(srv.ClerkClient, newUser.ClerkUserID, srv.Config)
	if err != nil {
		log.Debug().Err(err).Msg("NewClerkStore failed: Unable to create ClerkStore instance using clerk user id:")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Querying user_to_store table for user_to_store relationship for new employee
	newUTS, err := srv.DBClient.UserToStore.Query().
		Where(usertostore.StoreID(store.ID), usertostore.UserID(newUser.ID)).
		Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("UserJoinStore: unable to query from usertostore table in database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Adding the user_to_store relationship to Clerk's Data Store
	err = clerkStore.AddMetadata(newUTS)
	if err != nil {
		log.Debug().Err(err).Msg("UserJoinStore: unable to user metadata to Clerk Store")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully\n"))
}
