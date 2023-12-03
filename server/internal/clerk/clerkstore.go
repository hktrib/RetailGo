package clerkstorage

import (
	"encoding/json"
	"errors"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/rs/zerolog/log"
)

type ClerkStorage struct {
	client clerk.Client
	userID string
	user  *clerk.User
}


/*
addStoreToPublicMetadata:

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

				return &clerk.UpdateUserMetadata{
					PublicMetadata: bytes,
				}, nil
			}
		default:
			return nil, errors.New("unknown case type")
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
		return err
	}
	_, err = ch.client.Users().UpdateMetadata(ch.userID, publicMetadata)
	if err != nil {
		return err
	}
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