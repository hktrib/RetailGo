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
		fmt.Println("e.ID", e.ID)
		fmt.Println("event.Data.ID", event.Data.ID)
		if e.ID == event.Data.PrimaryEmailAddressID {
			return &e, nil
		}
	}
	return nil, errors.New("user not found")
}

func HandleClerkWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err).Msg("Unable to read from JSON request body")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

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
	svixHeaders.Set("svix-timestamp", svixTimestamp)
	svixHeaders.Set("svix-signature", svixSignature)

	// Creating New Webhook from CLERK_WEBHOOK_SECRET
	webhook, err := svix.NewWebhook(Config.CLERK_WEBHOOK_SECRET)
	if err != nil {
		log.Debug().Err(err).Msg("unable to create webhook")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Verifying payload against svixHeaders
	err = webhook.Verify([]byte(payload), svixHeaders)
	if err != nil {
		log.Debug().Err(err).Msg("webhook cannot be verified")
		http.Error(w, "Error verifying webhook", http.StatusBadRequest)
		return
	}

	// Process Event Data
	var event Event
	err = json.Unmarshal([]byte(payload), &event)
	fmt.Println(event.Data)
	if err != nil {
		log.Debug().Err(err).Msg("unable to martial json")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// user.created event handling
	if event.Type == UserCreated {
		emailObject, err := findUser(event)
		if err != nil {
			log.Debug().Err(err).Msg("failed finding user")
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

		url := "http://localhost:" + Config.SERVER_ADDRESS + "/create/user"
		// Use ngrok to create tempUrl for local testing purposes (io.ReadAll error handling sucks! with ngrok testing)
		// tempURl := "http://10a8-2600-1700-87f4-100-5125-57cc-354-6945.ngrok.io/create/user"
		b, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))

		client := &http.Client{}

		req.Header.Set("Authorization", "Bearer "+Config.CLERK_SK)

		resp, err := client.Do(req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
			return
		}
		defer resp.Body.Close()

		responseBody, _ := io.ReadAll(resp.Body)
		if err != nil {
			log.Debug().Err(err).Msg("failed to read response body")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		log.Debug().Msg(fmt.Sprintf("Response Code: %v, Response body: %v", resp.StatusCode, responseBody))
		w.WriteHeader(resp.StatusCode)
		w.Write(responseBody)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("error: event not supported!")
}
