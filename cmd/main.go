package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/NoCapCbas/webStash/internal/common"
	"github.com/NoCapCbas/webStash/internal/db/models"
	"github.com/NoCapCbas/webStash/internal/db/postgres/repos"
	"github.com/NoCapCbas/webStash/internal/db/sqlite"
	"github.com/NoCapCbas/webStash/internal/services"
)

type PageData struct {
	Title   string
	Message string
}

type LoginRequest struct {
	Email string `json:"email"`
}

var (
	templates = template.Must(
		template.New("").Funcs(common.FuncMap).ParseFiles(
			"templates/bookmarks/index.html",
			"templates/bookmarks/navbar.html",
			"templates/bookmarks/view.html",
			"templates/bookmarks/add-bookmark.html",
		),
	)
	authService     *services.AuthService
	mailgunService  services.MailService
	bookmarkService *services.BookmarkService
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// main portfolio handler
	fmt.Println("index...")

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
	magicLink, err := authService.GenerateMagicLink(req.Email)
	if err != nil {
		http.Error(w, "Error generating magic link", http.StatusInternalServerError)
		return
	}

	// Create user if doesn't exist
	log.Println("Creating user if doesn't exist")
	if err := authService.CreateUser(req.Email); err != nil {
		log.Println("Error creating user", err)
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Generate the full magic link URL
	log.Println("Generating full magic link URL")
	magicLinkURL := fmt.Sprintf("http://localhost:8080/verify?token=%s", magicLink.Token)

	// Send the magic link email
	log.Println("Sending magic link email", magicLinkURL, mailgunService)
	err = mailgunService.SendMagicLink(req.Email, magicLinkURL)
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
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Missing token", http.StatusBadRequest)
		return
	}

	// Validate the magic link token
	email, err := authService.ValidateMagicLink(token)
	if err != nil {
		log.Printf("Invalid magic link token: %v", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Generate a new session token instead of reusing the magic link token
	sessionToken, err := authService.GenerateSessionToken(email)
	if err != nil {
		log.Printf("Error generating session token: %v", err)
		http.Error(w, "Error creating session", http.StatusInternalServerError)
		return
	}

	// Set session cookie with the new token
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true, // Set to true if using HTTPS
		MaxAge:   60 * 60, // 1 hour
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/view/bookmarks", http.StatusSeeOther)
}

func bookmarkViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bookmark view handler")

	sessionTokenCookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println("No session token cookie found: ", err)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	// Use a separate function to validate sessions
	email, err := authService.ValidateSession(sessionTokenCookie.Value)
	if err != nil {
		log.Println("Invalid session token: ", err)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	bookmarks, err := bookmarkService.GetBookmarksByUserEmail(email)
	if err != nil {
		log.Println("Error getting bookmarks: ", err)
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}
	fmt.Println(bookmarks)
	// Render the account template with the email
	data := struct {
		Email     string
		Bookmarks []models.Bookmark
	}{
		Email:     email,
		Bookmarks: bookmarks,
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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/404.html")
	if err != nil {
		log.Println("Error parsing templates: ", err)
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func bookmarkCreateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bookmark create handler")
	var bookmark repos.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	bookmarkService.CreateBookmark(&bookmark)
}

func bookmarkUpdateHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bookmark update handler")
	var bookmark repos.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	bookmarkService.UpdateBookmark(&bookmark)
}

func bookmarkDeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bookmark delete handler")
	var bookmark repos.Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	bookmarkService.DeleteBookmark(bookmark.ID, bookmark.UserID)
}

type readBookmarkRequest struct {
	ID int `json:"id"`
}

func bookmarkReadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Bookmark read handler")
	var req readBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	bookmark, err := bookmarkService.GetBookmarkByID(req.ID)
	if err != nil {
		http.Error(w, "Bookmark not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(bookmark)
}

func main() {
	// Initialize database
	var err error
	sqlite, err := sqlite.New(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize repositories
	bookmarkRepo := repos.NewBookmarkRepo(sqlite.DB)
	userRepo := repos.NewUserRepo(sqlite.DB)
	sessionRepo := repos.NewSessionRepo(sqlite.DB)
	magicLinkRepo := repos.NewMagicLinkRepo(sqlite.DB)

	// Initialize services
	authService = services.NewAuthService(magicLinkRepo, sessionRepo, userRepo)
	mailgunService = services.NewMailgunService()
	bookmarkService = services.NewBookmarkService(bookmarkRepo)

	// serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// pass porfolio handler to server
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/verify", verifyHandler)
	http.HandleFunc("/view/bookmarks", bookmarkViewHandler)
	http.HandleFunc("/policies", policiesHandler)
	http.HandleFunc("/404", notFoundHandler)

	// bookmark handlers
	http.HandleFunc("/api/v1/bookmarks/create", bookmarkCreateHandler)
	http.HandleFunc("/api/v1/bookmarks/update", bookmarkUpdateHandler)
	http.HandleFunc("/api/v1/bookmarks/delete", bookmarkDeleteHandler)
	http.HandleFunc("/api/v1/bookmarks/read", bookmarkReadHandler)

	// start server
	log.Println("Listening on port :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
