package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/hktrib/RetailGo/internal/ent"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
)

func (srv *Server) GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Parse the store id
	store_id, err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if err != nil {
		fmt.Println("Error with Store Id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// get all employees belonging to this store
	user_data, err := srv.DBClient.User.Query().Where(user.StoreID(store_id)).All(ctx)

	if err != nil {
		fmt.Println("Issue with querying the DB.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user_bytes, err := json.Marshal(user_data)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(user_bytes)
}

// EmailRequest represents the request body for email testing
type EmailRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type HtmlTemplate struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Store_name  string `json:"store_name"`
	Sender_name string `json:"sender_name"`
	Action_url  string `json:"action_url"`
}

func (srv *Server) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	// Verify user credentials using clerk
	_, clerk_err := VerifyUserCredentials(r)
	if clerk_err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// get item id from url query string
	ctx := r.Context()
	user_id, id_err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if id_err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err := srv.DBClient.User.Query().Where(user.ID(user_id)).Only(ctx)

	if ent.IsNotFound(err) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

// TestEmailHandler handles the email testing request
func (srv *Server) SendInviteEmail(w http.ResponseWriter, r *http.Request) {

	entries, dummy_err := os.ReadDir("./")
	if dummy_err != nil {
		log.Fatal(dummy_err)
		fmt.Println(dummy_err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
	}

	ctx := r.Context()

	// Verify user credentials using clerk
	clerk_user, clerk_err := VerifyUserCredentials(r)
	if clerk_err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the store id
	store_id, id_err := strconv.Atoi(chi.URLParam(r, "store_id"))

	if id_err != nil {
		fmt.Println("Error with Store Id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Println("Store ID: ", store_id)
	storeObj, query_err := srv.DBClient.Store.Query().Where(store.ID(store_id)).Only(ctx)
	if query_err != nil {
		fmt.Println("Issue with querying the DB: " + query_err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"retailgoco@gmail.com",
		"hhqm pqgw lxmi weqb",
		"smtp.gmail.com",
	)

	// Define email headers and HTML content
	from := mail.Address{Name: "RetailGo Team", Address: "retailgoco@gmail.com"}
	to := mail.Address{Name: "Billy Bob", Address: req.Email}
	subject := "You have been invited to join a store!"

	// Message to be sent
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = `text/html; charset="UTF-8"`

	// HTML message

	tmpl, err_file := template.ParseGlob("./cmd/templates/*")
	if err_file != nil {
		// Handle error (e.g., file not found)
		http.Error(w, "Failed to open email template folder: "+err_file.Error(), http.StatusInternalServerError)
		return
	}

	var firstName, lastName string
	if clerk_user.FirstName != nil {
		firstName = *clerk_user.FirstName
	} else {
		// Handle the nil case for FirstName
		firstName = "unknown"
	}

	if clerk_user.LastName != nil {
		lastName = *clerk_user.LastName
	} else {
		// Handle the nil case for LastName
		lastName = "unknown"
	}

	// Read file content
	htmlBody := new(bytes.Buffer)
	templateData := HtmlTemplate{
		Email:       req.Email,
		Name:        req.Name,
		Store_name:  storeObj.StoreName,
		Sender_name: firstName + " " + lastName,
		//Sender_name: "Billy Bob",
		   Action_url: "http://localhost:3000/invite?code=" + storeObj.UUID,
		// Action_url: "https://retail-go.vercel.app/invite?code=" + storeObj.UUID,
	}
	err_io := tmpl.ExecuteTemplate(htmlBody, "email_invitation.html", templateData)

	if err_io != nil {
		// Handle error (e.g., error while reading)
		http.Error(w, "Failed to read email template: "+err_io.Error(), http.StatusInternalServerError)
		return
	}

	// Combine headers and HTML content
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + htmlBody.String()

	// Convert to and from to slice of strings
	toAddresses := []string{to.Address}

	// Send the email
	err := smtp.SendMail("smtp.gmail.com:587", auth, from.Address, toAddresses, []byte(message))
	if err != nil {
		http.Error(w, "Failed to send email: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Email sent successfully"))
}
