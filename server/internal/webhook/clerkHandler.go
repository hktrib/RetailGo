package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/hktrib/RetailGo/internal/util"
	"github.com/rs/zerolog/log"
	svix "github.com/svix/svix-webhooks/go"
)

var Config *util.Config
var ClerkClient clerk.Client

type EventType string

type EmailObject struct {
	EmailAddress string `json:"email_address"`
	ID           string `json:"id"`
}

type UserInterface struct {
	EmailObjects          []EmailObject `json:"email_addresses"`
	PrimaryEmailAddressID string        `json:"primary_email_address_id"`
	ID                    string        `json:"id"`
}

type Event struct {
	Data      UserInterface `json:"data"`
	Object    string        `json:"object"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Type      EventType     `json:"type"`
}

const (
	UserCreated = "user.created"
	UserUpdated = "user.updated"
	AnyEvent    = "*"
)

func findUser(event Event) (*EmailObject, error) {
	for _, e := range event.Data.EmailObjects {
		if e.ID == event.Data.ID {
			return &e, nil
		}
	}
	return nil, errors.New("user not found")
}

func HandleClerkWebhook(w http.ResponseWriter, r *http.Request) {
	// var payload map[string]interface{}
	// if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
	// 	http.Error(w, "Error decoding JSON", http.StatusBadRequest)
	// 	return
	// }

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err).Msg("Unable to read from JSON request body")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// payloadStr := fmt.Sprintf("%v", payload)

	svixID := r.Header.Get("svix-id")
	svixTimestamp := r.Header.Get("svix-timestamp")
	svixSignature := r.Header.Get("svix-signature")

	// Return error if headers are missing
	if svixID == "" || svixTimestamp == "" || svixSignature == "" {
		http.Error(w, "Error occurred", http.StatusBadRequest)
		return
	}

	// Setting Svix Http Headers for Verification
	svixHeaders := http.Header{}
	svixHeaders.Set("svix-id", svixID)
	svixHeaders.Set("svix-timestamp", svixSignature)
	svixHeaders.Set("svix-signature", svixTimestamp)
	// }

	// Creating New Webhook from CLERK_WEBHOOK_SECRET
	webhook, err := svix.NewWebhook(Config.CLERK_WEBHOOK_SECRET)
	if err != nil {
		fmt.Println("Error creating svix Webhook:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Verifying payload against svixHeaders
	err = webhook.Verify([]byte(payload), svixHeaders)
	if err != nil {
		fmt.Println("Error verifying webhook:", err)
		http.Error(w, "Error verifying webhook", http.StatusBadRequest)
		return
	}

	// Process Event Data
	var event Event
	err = json.Unmarshal([]byte(payload), &event)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// user.created event handling
	if event.Type == UserCreated {
		emailObject, err := findUser(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		// Creating Request Body for user creation request
		requestBody := map[string]interface{}{
			"email":         emailObject.EmailAddress,
			"is_owner":      false,
			"clerk_user_id": event.Data.ID,
			"first_name":    event.FirstName,
			"last_name":     event.LastName,
		}

		/* 	Using Clerk's .Do which is a wrapper that adds the Authorization key
		to the request on http.Client.Do and returns the response
		*/
		b, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest("PUT", "http://localhost:"+Config.SERVER_ADDRESS+"/create/user", bytes.NewBuffer(b))
		resp, err := ClerkClient.Do(req, nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Write(responseBody)
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("error: event not supported!")
}
