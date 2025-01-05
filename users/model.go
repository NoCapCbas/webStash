// contains data models
package users

import (
	"time"
)

type User struct {
	ID        int       `json:"id" validate:"required"`
	Verified  bool      `json:"verified" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	ContactInfo ContactInfo `json:"contact_info" validate:"required"`
	Address     Address     `json:"address" validate:"required"`

	Preferences Preferences `json:"preferences" validate:"required"`
	Account     Account     `json:"account" validate:"required"`
}

type ContactInfo struct {
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type Preferences struct {
	Language string `json:"language"`
	Timezone string `json:"timezone"`
}

type Account struct {
	ID int `json:"id" validate:"required"`

	MaxUsers  int       `json:"max_users" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	UpdatedAt time.Time `json:"updated_at" validate:"required"`
}
