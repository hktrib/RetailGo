package clerkstorage

import (
	"fmt"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/rs/zerolog/log"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/util"
)

const (
	PrivateMetadataType = "PrivateMetadata"
	PublicMetadataType = "PublicMetadata"
	UnsafeMetadata = "UnsafeMetadata"
	usersurl = "users"
	baseURL = "https://api.clerk.com/v1"
)

type ClerkStorage struct {
	client clerk.Client
	userID string
	user  *clerk.User
	Config *util.Config
}


type storeData struct {
	StoreID int `json:"id"`
	Name string `json:"storename"`
	PermissionLevel int `json:"Permission_level"`
}

func Equal(a storeData, b storeData) bool {
	return a.StoreID == b.StoreID && a.Name == b.Name && a.PermissionLevel == b.PermissionLevel
}

func NewClerkStore(client clerk.Client, clerkUserID string, config *util.Config) (*ClerkStorage, error) {
	cus := &ClerkStorage{}
	cus.client = client
	cus.userID = clerkUserID
	cus.Config = config

	// Fetching User{} from Clerk Store
	user, err := client.Users().Read(clerkUserID) 
	if err != nil {
		return nil, err
	}

	cus.user = user
	return cus, nil
}


func (cs ClerkStorage) AddMetadata(userStoreRelation *ent.UserToStore) error {
	clerkUserMetadata, err := cs.addStoreToPublicMetadata(cs.user, userStoreRelation)	
	if err != nil {
		log.Debug().Err(err).Msg("Unable to Add Store to Public Metadata")
		return err
	}

	fmt.Printf("User ID: %v | public metadata: %v\n", cs.userID, *clerkUserMetadata)

	user, err := cs.updateMetadata(*clerkUserMetadata)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to update metadata")
		return err
	}
	
	fmt.Printf("User PublicMetadata: %v\n", user.PublicMetadata)
	return nil
}

func (cs ClerkStorage) RemoveMetadata(userStoreRelation *ent.UserToStore) error {
	
	clerkUserMetadata, err := cs.removeStoreFromPublicMetadata(cs.user, userStoreRelation)	
	if err != nil {
		log.Debug().Err(err).Msg("Unable to Remove Store from User's Public Metadata")
		return err
	}

	log.Debug().Msg(fmt.Sprintf("User ID: %v | public metadata: %v\n", cs.userID, *clerkUserMetadata))

	user, err := cs.updateMetadata(*clerkUserMetadata)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to update metadata")
		return err
	}
	
	log.Debug().Msg(fmt.Sprintf("User PublicMetadata: %v\n", user.PublicMetadata))
	return nil
}