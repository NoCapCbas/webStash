// contains data models
package users

import (
	"time"
)

type Admin bool

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Verified  bool      `json:"verified"`
	Admin     Admin     `json:"admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
