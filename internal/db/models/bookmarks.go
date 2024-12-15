package models

import "time"

type Bookmark struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	URL         string    `json:"url" db:"url"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Public      bool      `json:"public" db:"public"`
	ClickCount  int       `json:"click_count" db:"click_count"`
}
