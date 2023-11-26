package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/mail"
	"strconv"

	"net/smtp"

	"github.com/go-chi/chi/v5"
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

	// Verify valid store id (not SQL Injection)

	// Run the SQL Query by store id on the items table, to get all items belonging to this store
	user_data, err := srv.DBClient.User.Query().Where(user.StoreID(store_id)).All(ctx)

	if err != nil {
		fmt.Println("Issue with querying the DB.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	user_bytes, err := json.MarshalIndent(user_data, "", "")

	// Format and return

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)

	w.Write(user_bytes)
}

// EmailRequest represents the request body for email testing
type EmailRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Store_name  string `json:"store_name"`
	Sender_name string `json:"sender_name"`
}

type HtmlTemplate struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Store_name  string `json:"store_name"`
	Sender_name string `json:"sender_name"`
	Action_url  string `json:"action_url"`
}

// TestEmailHandler handles the email testing request
func (srv *Server) TestEmailHandler(w http.ResponseWriter, r *http.Request) {

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"retailgoco@gmail.com",
		"iirp tdgd sbcp ktsp",
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

	tmpl, err_file := template.ParseFiles("templates/email_invitation.html")
	if err_file != nil {
		// Handle error (e.g., file not found)
		http.Error(w, "Failed to open email template: "+err_file.Error(), http.StatusInternalServerError)
		return
	}

	// Read file content
	htmlBody := new(bytes.Buffer)
	Template := []HtmlTemplate{
		{
			Email:       req.Email,
			Name:        req.Name,
			Store_name:  req.Store_name,
			Sender_name: req.Sender_name,
			Action_url:  "http://localhost:300/sign-up?" + "1381",
		},
	}
	err_io := tmpl.Execute(htmlBody, Template)
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
