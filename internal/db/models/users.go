package models

import "time"

type User struct {
	ID             int       `json:"id" db:"id"`
	Email          string    `json:"email" db:"email"`
	MembershipType string    `json:"membership_type" db:"membership_type"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}