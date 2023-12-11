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

	"github.com/hktrib/RetailGo/internal/ent"
)

type ClerkUserMetadata struct {
	// Empty map to allow any type of data for public metadata
	PublicMetadata map[string]interface{} `json:"public_metadata"`

	// Empty map to allow any type of data for private metadata
	PrivateMetadata map[string]interface{} `json:"private_metadata"`

	// Empty map to allow any type of data for unsafe metadata
	UnsafeMetadata map[string]interface{} `json:"unsafe_metadata"`
}

func (cs ClerkStorage) NewClerkUserMetadata() (*ClerkUserMetadata, error) {
	cum := &ClerkUserMetadata{}
	switch cs.user.PublicMetadata.(type) {
	case map[string]interface{}:
		log.Debug().Msg("PublicMetadata Case: map[interface{}]interface{}")
		v, ok := cs.user.PublicMetadata.(map[string]interface{})
		if ok {
			cum.PublicMetadata = v
		} else {
			cum.PublicMetadata = make(map[string]interface{})
		}

	default:
		return nil, errors.New("something went wrong terribly wrong creating PublicMetadata")
	}

	switch cs.user.PrivateMetadata.(type) {
	case map[string]interface{}:
		log.Debug().Msg("PublicMetadata Case: map[interface{}]interface{}")
		v, ok := cs.user.PublicMetadata.(map[string]interface{})
		if ok {
			cum.PublicMetadata = v
		} else {
			cum.PublicMetadata = make(map[string]interface{})
		}

	default:
		return nil, errors.New("something went wrong terribly wrong creating PrivateMetadata")
	}

	switch cs.user.UnsafeMetadata.(type) {
	case map[string]interface{}:
		log.Debug().Msg("PublicMetadata Case: map[interface{}]interface{}")
		v, ok := cs.user.PublicMetadata.(map[string]interface{})
		if ok {
			cum.PublicMetadata = v
		} else {
			cum.PublicMetadata = make(map[string]interface{})
		}

	default:
		return nil, errors.New("something went wrong terribly wrong creating UnsafeMetadata")
	}
	return cum, nil
}

func (cs ClerkStorage) addStoreToPublicMetadata(user *clerk.User, uts *ent.UserToStore) (*ClerkUserMetadata, error) {

	metadata, err := cs.NewClerkUserMetadata()
	if err != nil {
		return nil, err
	}

	stores, ok := metadata.PublicMetadata["stores"].([]interface{})
	if !ok {
		metadata.PublicMetadata["stores"] = []storeData{{
			StoreID:         uts.StoreID,
			Name:            uts.StoreName,
			PermissionLevel: uts.PermissionLevel,
		}}
	} else {
		newStoreData := storeData{
			StoreID:         uts.StoreID,
			Name:            uts.StoreName,
			PermissionLevel: uts.PermissionLevel,
		}

		for _, data := range stores {
			store, ok := data.(storeData)
			if !ok {
				return nil, fmt.Errorf("malformed data, %v", data)
			}
			if Equal(store, newStoreData) {
				return metadata, nil
			}
		}
		metadata.PublicMetadata["stores"] = append(stores, newStoreData)
	}
	return metadata, nil
}

func (cs ClerkStorage) removeStoreFromPublicMetadata(user *clerk.User, uts *ent.UserToStore) (*ClerkUserMetadata, error) {

	metadata, err := cs.NewClerkUserMetadata()
	if err != nil {
		return nil, err
	}

	stores, ok := metadata.PublicMetadata["stores"].([]storeData)
	if !ok {
		return nil, fmt.Errorf("no Clerk Metadata! Something went wrong")
	} else {
		storeDataToDelete := storeData{
			StoreID:         uts.StoreID,
			Name:            uts.StoreName,
			PermissionLevel: uts.PermissionLevel,
		}

		storeIndexToDelete := -1
		for index, store := range stores {
			if Equal(store, storeDataToDelete) {
				storeIndexToDelete = index
				break
			}
		}

		if storeIndexToDelete != -1 {
			if storeIndexToDelete == 0 {
				metadata.PublicMetadata["stores"] = stores[:1]
			} else if storeIndexToDelete == len(stores)-1 {
				metadata.PublicMetadata["stores"] = stores[len(stores)-1:]
			} else {
				metadata.PublicMetadata["stores"] = append(stores[:storeIndexToDelete], stores[storeIndexToDelete+1:]...)
			}
		} else {
			log.Debug().Msg(fmt.Sprintf("Stores: %v | UserToStore", stores))
			return nil, fmt.Errorf("store index %d not found", storeIndexToDelete)
		}
	}

	return metadata, nil
}

func (cs ClerkStorage) updateMetadata(c ClerkUserMetadata) (*clerk.User, error) {
	url := fmt.Sprintf("%s/%s/%s/metadata", baseURL, usersurl, cs.userID)

	b, _ := json.Marshal(c)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cs.Config.CLERK_SK)

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
