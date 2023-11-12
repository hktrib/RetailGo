package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/ent"
	store2 "github.com/hktrib/RetailGo/ent/store"
	user2 "github.com/hktrib/RetailGo/ent/user"
	"github.com/hktrib/RetailGo/transactions"
	"github.com/rs/zerolog/log"
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
			log.Fatal().Err(err).Msg("failed to identify user by email")
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

func (srv *Server) OwnerCreate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		// user := ctx.Value("user").(*ent.User)

		// Checking if the user is an owner
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		potentialOwner := ent.User{}

		err = json.Unmarshal(reqBody, &potentialOwner)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if potentialOwner.StoreID == 0 {
			taskLifeTime := 1 * time.Minute
			err := srv.TaskProducer.TaskOwnerCreationCheck(ctx, &potentialOwner.Email, taskLifeTime)
			if err != nil {
				log.Fatal().Err(err).Msg("TaskOwnerCreationCheck failed..")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			srv.Cache.SetX(potentialOwner.Email, &potentialOwner, taskLifeTime)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("OK!"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (srv *Server) StoreCreate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()

		type Request struct {
			Email     string `json:"email"`
			StoreName string `json:"name"`
		}

		reqBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal().Err(err).Msg("Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		reqBody := Request{}

		err = json.Unmarshal(reqBytes, &reqBody)
		if err != nil {
			log.Fatal().Err(err).Msg("Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		store := ent.Store{}
		store.StoreName = reqBody.StoreName
		owner := srv.Cache.Get(reqBody.Email)

		ctx = context.WithValue(ctx, transactions.StoreKey{Store: "store"}, store)
		_ = context.WithValue(ctx, transactions.OwnerKey{Owner: "owner"}, owner)

		next.ServeHTTP(w, r)
	})
}
