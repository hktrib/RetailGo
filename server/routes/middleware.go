package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent"
	store2 "github.com/hktrib/RetailGo/ent/store"
	user2 "github.com/hktrib/RetailGo/ent/user"
)

func (srv *Server) ValidateStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store_id := chi.URLParam(r, "store_id")
		StoreId, err := strconv.Atoi(store_id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		store, err := srv.DBClient.Store.
			Query().Where(store2.ID(StoreId)).Only(r.Context())
		if ent.IsNotFound(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), "store_var", store)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) GetAuthenticatedUserEmail(ctx context.Context) (string, error) {
	sessClaims, ok := clerk.SessionFromContext(ctx)
	if !ok {
		return "", errors.New("Error: Authentication Error -> No session claims")
	}

	user, err := srv.ClerkClient.Users().Read(sessClaims.Claims.Subject)
	email := user.EmailAddresses[0].EmailAddress
	if err != nil {
		return "", errors.New("Error: Authentication Error -> Unable to get User from Clerk")
	}
	return email, nil
}

func (srv *Server) ValidateStoreAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Getting the store id from the request URL
		param := chi.URLParam(r, "store_id")
		storeID, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Getting the emailAddress to identify user who made the request
		ctx := r.Context()
		emailAddress, err := srv.GetAuthenticatedUserEmail(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatal(err)
		}

		user, err := srv.DBClient.User.Query().Where(user2.Email(emailAddress)).Only(ctx)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Validating if the user who submitted the request has access to the store.
		if user.StoreID != storeID {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx = context.WithValue(ctx, "store_var", storeID)
		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}

func (srv *Server) ValidateOwner(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user := ctx.Value("user").(*ent.User)

		// Checking if the user is an owner
		if !user.IsOwner {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
