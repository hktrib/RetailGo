package clerkstorage

import (
	"encoding/json"
	"fmt"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/rs/zerolog/log"
)

type ClerkStorage struct {
	client clerk.Client
	userID string
	user  *clerk.User
}


/*
Accepts (clerk.User, int (store id))
________________________________________

-> It adds storeID to "stores" field of PublicMetadata

-> Creates the "stores" field if it does not exist it 

External Package Calls:
	json.Marshal
*/
func addStoreToPublicMetadata(user *clerk.User, storeID int) (*clerk.UpdateUserMetadata, error) {
	switch user.PublicMetadata.(type){
		case map[any]any:
			log.Debug().Msg("PrivateMetadata Case: map[interface{}]interface{}")
			v := user.PublicMetadata.(map[string]any)
			list, ok := v["stores"].([]int)
			if ok {
				list := append(list, storeID)
				stores := map[string]any{
					"stores": list,
				}

				bytes, err := json.Marshal(stores)
				if err != nil {
					return nil, err
				}

				return &clerk.UpdateUserMetadata{PublicMetadata: bytes}, nil
			}

		case any:
			stores := map[string][]int{
				"stores": {storeID},
			}

			bytes, err := json.Marshal(stores)
			if err != nil {
				return nil, err
			}

			log.Debug().Msg("No Previous data in PublicMetadata")

			return &clerk.UpdateUserMetadata{PublicMetadata: bytes}, nil
		default:
			log.Debug().Msg("Public Metadata diff type not supported")
	}
	return nil, nil
}


/*
Accepts (int (denoting store id))

-----------------------------------------
		
	Updates the Clerk User Metadata
	
	It returns an error if it fails and nil otherwise.

External Package Calls:
	
	-> clerk.UsersService.UpdateUserMetadata
*/
func (ch ClerkStorage) AddStore(storeID int) error {
	publicMetadata, err := addStoreToPublicMetadata(ch.user, storeID)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to Add Store to Public Metadata")
		return err
	}
	user, err := ch.client.Users().UpdateMetadata(ch.userID, publicMetadata)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to update metadata")
		return err
	}
	
	fmt.Printf("User PublicMetadata: %v\n", user.PublicMetadata)
	return nil
}


/*

NewClerkStore 

-> Accepts clerk.Client and a string representing the clerks user id
// It returns a new ClerkStore instance

*/
func NewClerkStore(client clerk.Client, clerkUserID string) (*ClerkStorage, error) {
	cus := &ClerkStorage{}
	cus.client = client
	cus.userID = clerkUserID

	// Fetching User{} from Clerk Store
	user, err := client.Users().Read(clerkUserID) 
	if err != nil {
		return nil, err
	}

	cus.user = user
	return cus, nil
}