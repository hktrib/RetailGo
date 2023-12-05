package clerkstorage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/rs/zerolog/log"

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


func (cs ClerkStorage) addStoreToPublicMetadata(user *clerk.User, storeID int) (*ClerkUserMetadata, error) {
	
	metadata, err := cs.NewClerkUserMetadata()
	if err != nil {
		return nil, err
	}

	stores, ok := metadata.PublicMetadata["stores"].([]interface{})
	if !ok {
		metadata.PublicMetadata["stores"] = []int{storeID}
	} else {
		stores = append(stores, storeID)
		metadata.PublicMetadata["stores"] = stores
	}
	return metadata, nil
}


func (cs ClerkStorage) updateMetadata(c ClerkUserMetadata) (*clerk.User, error) {
	url := fmt.Sprintf("%s/%s/%s/metadata", baseURL, usersurl, cs.userID)

	b, _ := json.Marshal(c)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ cs.Config.CLERK_SK)


	var user clerk.User

	client := &http.Client{}
	
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to send request")
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)
	if err != nil {
		log.Debug().Err(err).Msg("failed to read response body")
		return nil, errors.New("failed to read response body")
	}

	err = json.Unmarshal(responseBody, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body | StatusCode %v", resp.StatusCode)
	}

	return &user, nil
}

func (cs ClerkStorage) AddStore(storeID int) error {
	clerkUserMetadata, err := cs.addStoreToPublicMetadata(cs.user, storeID)	
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

	// user, err := ch.client.Users().UpdateMetadata(ch.userID, publicMetadata)
	// if err != nil {
	// 	log.Debug().Err(err).Msg("Unable to update metadata")
	// 	return err
	// }
	
	fmt.Printf("User PublicMetadata: %v\n", user.PublicMetadata)
	return nil
}

