package server

import (
	"encoding/json"
	"fmt"
	. "github.com/hktrib/RetailGo/internal/stripe-components"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	clerkstorage "github.com/hktrib/RetailGo/internal/clerk"
	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
	"github.com/hktrib/RetailGo/internal/ent/usertostore"
	"github.com/hktrib/RetailGo/internal/transactions"
	"github.com/rs/zerolog/log"
)

/*
CreateStore Brief:

-> Creates a new store and its owner, utilizing transactions to ensure atomicity.

	Creates a ClerkStore instance, then executes a transaction to create the store and its owner.

External Package Calls:
- clerkstorage.NewClerkStore()
- transactions.StoreAndOwnerCreationTx()
*/
func (srv *Server) CreateStore(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	reqStore := ctx.Value(Param("store")).(*ent.Store)
	reqUser := ctx.Value(Param("owner")).(*ent.User)

	clerkStore, err := clerkstorage.NewClerkStore(srv.ClerkClient, reqUser.ClerkUserID, srv.Config)
	if err != nil {
		log.Debug().Err(err).Msg("CreateStore: NewClerkStore failed -> Unable to create ClerkStore instance using clerk user id:")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err, nStore, nUser := transactions.StoreAndOwnerCreationTx(ctx, reqStore, reqUser, srv.DBClient, clerkStore)
	err = SendOnboardingEmail(nStore, nUser)
	if err != nil {
		log.Debug().Err(err).Msg("CreateStore: unable to send onboarding email")
		return
	}

	if err != nil {
		log.Debug().Err(err).Msg("CreateStore: could not executed Store|Owner Transaction ")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err != nil {
		log.Debug().Err(err).Msg("CreateStore: unable to send onboarding email")
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(nStore.ID)))
}

/*
GetStoreUsers Brief:

-> Retrieves users associated with a specific store by Store ID.

	Fetches users related to the store and prepares a response containing relevant user information.

External Package Calls:
- srv.DBClient.User.Query()
*/
func (srv *Server) GetStoreUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	StoreID, err := strconv.Atoi(chi.URLParam(r, "store_id"))
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreUsers: unable to parse store_id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	Users, err := srv.DBClient.User.Query().Where(user.StoreID(StoreID)).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreUsers: server unable to fetch all store's users from database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(Users)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreUsers: server unable to Marshal store's users")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var userInfo []map[string]interface{}
	err = json.Unmarshal(resp, &userInfo)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreUsers: server unable to Unmarshal store's users")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for i := range userInfo {
		delete(userInfo[i], "clerk_user_id")
		delete(userInfo[i], "store_id")
		delete(userInfo[i], "edges")

	}

	resp, err = json.Marshal(userInfo)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreUsers: server unable to Marshal store's userInfo response")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
func (srv *Server) GetStoreByUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uuid := chi.URLParam(r, "uuid")
	// Query the store by UUID
	store, err := srv.DBClient.Store.Query().Where(store.UUID(uuid)).Only(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreByUUID: server unable to fetch the store from database")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Marshalling the store data
	resp, err := json.Marshal(PruneStore(store))
	if err != nil {
		log.Debug().Err(err).Msg("GetStoreByUUID: server unable to Marshal the store")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Writing the response
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

/*
GetStoresByClerkId Brief:

-> Retrieves stores associated with a specific clerk by Clerk user ID.
*/
func (srv *Server) GetStoresByClerkId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Verify user credentials using clerk
	clerkUser, clerkErr := VerifyUserCredentials(r)
	if clerkErr != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+clerkErr.Error(), http.StatusInternalServerError)
		w.Write([]byte("Unauthorized"))
		return
	}

	// Query the store by UUID
	stores, err := srv.DBClient.UserToStore.Query().Where(usertostore.ClerkUserID(clerkUser.ID)).All(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("GetStoresByClerkId: server unable to fetch the stores from database")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Marshalling the store data
	resp, err := json.Marshal(PruneUserStore(stores...))
	if err != nil {
		log.Debug().Err(err).Msg("GetStoresByClerkId: server unable to Marshal the store")
		http.Error(w, http.StatusText(http.StatusInternalServerError)+": "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Writing the response
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (srv *Server) HandleOnboarding(w http.ResponseWriter, r *http.Request) {
	store_id, _ := strconv.Atoi(chi.URLParam(r, "store_id"))

	// Get store
	store, err := srv.DBClient.Store.Query().Where(store.ID(store_id)).Only(r.Context())
	if err != nil {
		log.Debug().Err(err).Msg("HandleOnboarding: unable to fetch store from database")
		return
	}
	accLink, err := StartOnboarding(store.StripeAccountID, store_id)
	if err != nil {
		log.Debug().Err(err).Msg("HandleOnboarding: unable to start onboarding")
		return
	}
	w.WriteHeader(302)
	fmt.Fprintf(w, "%s", accLink.URL)
}

func (srv *Server) GetAuthorizationStatus(w http.ResponseWriter, r *http.Request) {
	store_id, _ := strconv.Atoi(chi.URLParam(r, "store_id"))

	// Get store
	store, err := srv.DBClient.Store.Query().Where(store.ID(store_id)).Only(r.Context())
	if err != nil {
		log.Debug().Err(err).Msg("GetAuthorizationStatus: unable to fetch store from database")
		return
	}
	if store.IsAuthorized {
		w.WriteHeader(200)
		w.Write([]byte("true"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("false"))
}
