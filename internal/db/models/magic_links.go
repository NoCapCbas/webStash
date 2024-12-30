package models

import "time"

type MagicLink struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Token     string    `json:"token" db:"token"`
	Used      bool      `json:"used" db:"used"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
