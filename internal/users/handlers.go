package users

import (
	"net/http"
	"encoding/json"
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


