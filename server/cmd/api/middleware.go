package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/hktrib/RetailGo/internal/ent"
	store2 "github.com/hktrib/RetailGo/internal/ent/store"
	user2 "github.com/hktrib/RetailGo/internal/ent/user"
	"github.com/rs/zerolog/log"
)

func (srv *Server) ValidateStore(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store_id := chi.URLParam(r, "store_id")
		StoreId, err := strconv.Atoi(store_id)
		if err != nil {
			log.Debug().Err(err).Msg(fmt.Sprintf("store id %s is not a integer", store_id))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		store, err := srv.DBClient.Store.
			Query().Where(store2.ID(StoreId)).Only(r.Context())
		if ent.IsNotFound(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		} else if err != nil {
			log.Debug().Err(err).Msg(fmt.Sprintf("duplicate store on store id %s", store))
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
		return "", errors.New("error: Authentication Error -> No session claims")
	}

	user, err := srv.ClerkClient.Users().Read(sessClaims.Claims.Subject)

	// Maybe we need to add another check here to validate email existence.
	email := user.EmailAddresses[0].EmailAddress
	if err != nil {
		return "", errors.New("error: Authentication Error -> Unable to get User from Clerk")
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
			log.Debug().Err(err).Msg(fmt.Sprintf("requested store_id %s not a integer", param))
			return
		}

		// Getting the emailAddress to identify user who made the request
		ctx := r.Context()
		emailAddress, err := srv.GetAuthenticatedUserEmail(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Debug().Err(err).Msg("failed to identify user by email")
			return
		}

		user, err := srv.DBClient.User.Query().Where(user2.Email(emailAddress)).Only(ctx)
		if err != nil {
			log.Debug().Err(err).Msg("unable to find user by email")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Validating if the user who submitted the request has access to the store.
		if user.StoreID != storeID {
			log.Debug().Msg(fmt.Sprintf("%v: user doesnt have access to store!!!", user))
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx = context.WithValue(ctx, "store_var", storeID)
		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) ValidateOwner(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user := ctx.Value("user").(*ent.User)

		// Checking if the user is an owner
		if !user.IsOwner {
			log.Debug().Msg(fmt.Sprintf("%v: user isn't owner!!!", user))
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (srv *Server) IsOwnerCreateHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// user := ctx.Value("user").(*ent.User)

		// Checking if the user is an owner
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Debug().Msg("Unable to unmarshal")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		potentialOwner := ent.User{}

		err = json.Unmarshal(reqBody, &potentialOwner)
		if err != nil {
			log.Debug().Msg("Unable to unmarshal")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if potentialOwner.StoreID == 0 {
			taskLifeTime := 1 * time.Minute
			err := srv.TaskProducer.TaskOwnerCreationCheck(ctx, &potentialOwner.Email, taskLifeTime)
			if err != nil {
				log.Debug().Err(err).Msg("TaskOwnerCreationCheck failed..")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			srv.Cache.SetX(potentialOwner.Email, &potentialOwner, 10*taskLifeTime)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("OK!"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (srv *Server) StoreCreateHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		reqBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Debug().Err(err).Msg("Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		store := ent.Store{}

		err = json.Unmarshal(reqBytes, &store)
		if err != nil {
			log.Debug().Err(err).Msg("Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		owner := srv.Cache.Get(store.OwnerEmail)
		if owner == nil {
			log.Debug().Err(err).Msg("Owner's user data is outdated...")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx = context.WithValue(ctx, "store", &store)
		ctx = context.WithValue(ctx, "owner", owner)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (srv *Server) ProtectStoreAndOwnerCreation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		clerkUserEmailAddress, err := srv.GetAuthenticatedUserEmail(ctx)
		if err != nil {
			log.Err(err).Msg("failed to identify user by email")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if clerkUserEmailAddress == "" {
			log.Err(err).Msg("failed to identify user by email")
			http.Error(w, errors.New("no email associated with the session claims").Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
