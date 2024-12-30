// contains data models
package users

import (
	"time"
)

type Role int

const (
	Admin Role = iota // Starts at 0 and increments
	Basic	
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Verified bool `json:"verified"`
	Role     Role  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`	
}

// User Specific Events
const (
	UserSignedUpEvent = "user.signedup"
	UserUpdatedEvent = "user.updated"
	UserVerifiedEvent = "user.verified"
)
