package users

import (
	"log"	
	"net/http"
	"github.com/NoCapCbas/webStash/internal/users"
)

func main() {
	
	// Initialize the service
	userService := NewUserService()

	// Initialize the handler
	userHandler := NewUserHandler(userService)

	// Set up general user routes /{service}/{event}
	http.HandleFunc("/users/signup", SignUpUserHandler)

	// Set up user specific routes /{service}/{event}/{user_id}
	http.HandleFunc("users/login/{id}", userHandler.LoginUserHandler)
	http.HandleFunc("users/update/{id}", userHandler.UpdateUserHandler)
	http.HandleFunc("users/deactivate/{id}", userHandler.DeactivateUserHandler)
	http.HandleFunc("users/reactivate/{id}", userHandler.ReactivateUserHandler)

  // Start the server
	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)	
	if err != nil {
		log.Fatal(err)
	}
}
