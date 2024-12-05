package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/NoCapCbas/webStash/internal/auth"
	"github.com/NoCapCbas/webStash/internal/common"
	"github.com/NoCapCbas/webStash/internal/db"
)

type PageData struct {
	Title   string
	Message string
}

type LoginRequest struct {
	Email string `json:"email"`
}

var postgres *db.PostgresDB
var templates = template.Must(template.ParseFiles("templates/account/index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// main portfolio handler

	// parse template file
	tmpl, err := template.ParseFiles(
		"templates/index.html",
		"templates/partials/nav.html",
		"templates/partials/hero.html",
		"templates/partials/features.html",
		"templates/partials/pricing-1-plan.html",
		"templates/partials/footer-cta.html",
		"templates/partials/footer.html",
	)
	if err != nil {
		log.Println("Error parsing templates: ", err)
		http.Error(w, "Could not load template: ", http.StatusInternalServerError)
		return
	}

	// create some data to pass to the template
	data := PageData{
		Title:   "title",
		Message: "message",
	}

	// execute the template and pass the data
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Generate magic link
	log.Println("Generating magic link for", req.Email)
	magicLink, err := auth.GenerateMagicLink(req.Email)
	if err != nil {
		http.Error(w, "Error generating magic link", http.StatusInternalServerError)
		return
	}

	// Create user if doesn't exist
	log.Println("Creating user if doesn't exist")
	if err := postgres.CreateUser(req.Email); err != nil {
		log.Println("Error creating user", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Initialize Mailgun client
	mailgunClient := common.NewMailgunClient()

	// Generate the full magic link URL
	log.Println("Generating full magic link URL")
	magicLinkURL := fmt.Sprintf("http://localhost:8080/verify?token=%s", magicLink.Token)

	// Send the magic link email
	log.Println("Sending magic link email", magicLinkURL, mailgunClient)
	err = mailgunClient.SendMagicLink(req.Email, magicLinkURL)
	if err != nil {
		log.Printf("Failed to send magic link: %v", err)
		http.Error(w, "Error sending magic link", http.StatusInternalServerError)
		return
	}

	log.Println("Magic link sent to email")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Magic link sent to email",
	})
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Verifying magic link")
	// Get token from URL query parameters
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	// Validate the magic link token
	_, err := auth.ValidateMagicLink(token)
	if err != nil {
		log.Printf("Invalid magic link token: %v", err)
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Set to true if using HTTPS
		MaxAge:   60 * 15, // 15 minutes
		SameSite: http.SameSiteStrictMode,
	})

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}

func accountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Account handler")

	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	// Validate the magic link token
	email, err := auth.ValidateMagicLink(sessionTokenCookie.Value)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	// Render the account template with the email
	data := struct {
		Email string
	}{
		Email: email,
	}

	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func policiesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"templates/policies/main.html",
		"templates/partials/nav.html",
		"templates/partials/footer.html",
	)
	if err != nil {
		log.Println("Error parsing templates: ", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template: ", err)
		http.Error(w, "Could not render template", http.StatusInternalServerError)
	}
}

func main() {
	// Initialize database
	var err error
	postgres, err = db.NewPostgresDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// pass porfolio handler to server
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/verify", verifyHandler)
	http.HandleFunc("/account", accountHandler)
	http.HandleFunc("/policies", policiesHandler)

	// start server
	log.Println("Listening on port :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
