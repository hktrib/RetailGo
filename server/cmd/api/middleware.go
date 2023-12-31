package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/hktrib/RetailGo/internal/ent"
	user2 "github.com/hktrib/RetailGo/internal/ent/user"
)

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
			log.Debug().Err(err).Msg(fmt.Sprintf("requested store_id %v not a integer", storeID))
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

		ctx = context.WithValue(ctx, Param("store_var"), storeID)
		ctx = context.WithValue(ctx, Param("user"), user)

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
func (srv *Server) StoreExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		storeId, err := strconv.Atoi(chi.URLParam(r, "store_id"))
		if err != nil {
			log.Debug().Err(err).Msg("unable to parse store_id")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		store, err := srv.DBClient.Store.Query().Where(store.ID(storeId)).Only(ctx)
		if err != nil {
			log.Debug().Err(err).Msg("unable to find store by id")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx = context.WithValue(ctx, Param("store"), store)
		next.ServeHTTP(w, r.WithContext(ctx))
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

		fmt.Printf("PotentialOwner %v\n", potentialOwner)

		if potentialOwner.StoreID == 0 {
			fmt.Println("HIT BRANCH")
			taskLifeTime := 10 * time.Minute
			err := srv.TaskProducer.ProduceTaskOwnerCreationCheck(ctx, &potentialOwner.Email, taskLifeTime)
			if err != nil {
				log.Debug().Err(err).Msg("TaskOwnerCreationCheck failed..")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			srv.Cache.Set("storecreate|"+potentialOwner.Email, &potentialOwner)
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
			log.Debug().Err(err).Msg("io.Readall: Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		store := ent.Store{}

		err = json.Unmarshal(reqBytes, &store)
		if err != nil {
			log.Debug().Err(err).Msg("json.Unmarshal: Reading request body failed..")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		owner := srv.Cache.Get("storecreate|" + store.OwnerEmail)
		if owner == nil {
			log.Debug().Err(err).Msg("Owner's user data is outdated...")
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		ctx = context.WithValue(ctx, Param("store"), &store)
		ctx = context.WithValue(ctx, Param("owner"), owner)

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

// VerifyUserCredentials verifies the user's credentials and returns user data if valid.
func VerifyUserCredentials(r *http.Request) (*clerk.User, error) {
	client, err := clerk.NewClient(os.Getenv("CLERK_SK"))
	if err != nil {
		// Handle error in creating Clerk client
		return nil, err
	}

	// Get the session token from the Authorization header
	sessionToken := r.Header.Get("Authorization")
	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")

	// Verify the session
	sessClaims, err := client.VerifyToken(sessionToken)
	if err != nil {
		// Handle error in verifying the token
		return nil, err
	}

	// Get the user, and say welcome!
	user, err := client.Users().Read(sessClaims.Claims.Subject)
	if err != nil {
		// Handle error in reading user data
		return nil, err
	}

	return user, nil
}
