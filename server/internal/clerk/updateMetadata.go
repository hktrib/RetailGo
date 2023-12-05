package clerkstorage

import (
	"errors"

	"github.com/rs/zerolog/log"
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
		v, ok:= cs.user.PublicMetadata.(map[string]interface{})
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
