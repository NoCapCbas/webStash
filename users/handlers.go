package users

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) SignUpUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set the header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a map with the key-value pair you want in the JSON response
	response := map[string]string{"message": "Hello, World!"}

	// Encode the map into JSON and send it in the response
	w.WriteHeader(http.StatusOK) // Status 200 OK
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	// Set the header to indicate that the response is JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a map with the key-value pair you want in the JSON response
	response := map[string]string{"message": "Hello, World!"}

	// Encode the map into JSON and send it in the response
	w.WriteHeader(http.StatusOK) // Status 200 OK
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandler) DeactivateUserHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandler) ReactivateUserHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func (h *UserHandler) EventTypesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a map of event types and their descriptions
	eventTypes := map[string]string{
		string(UserCreatedEvent):  "Triggered when a new user is created",
		string(UserVerifiedEvent): "Triggered when a user is verified",
		string(UserUpdatedEvent):  "Triggered when user information is updated",
		// Add any other event types here
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eventTypes); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
