package repos

import (
	"database/sql"
	"log"
	"time"
)

type MagicLink struct {
	ID        int
	Email     string
	Token     string
	Used      bool
	ExpiresAt time.Time
	CreatedAt time.Time
}

type MagicLinkRepo struct {
	db *sql.DB
}

func NewMagicLinkRepo(db *sql.DB) *MagicLinkRepo {
	return &MagicLinkRepo{db: db}
}

func (r *MagicLinkRepo) Create(email string, token string, expiresAt time.Time) (*MagicLink, error) {
	var ml MagicLink
	err := r.db.QueryRow(`
		INSERT INTO magic_links (email, token, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, email, token, used, expires_at, created_at
	`, email, token, expiresAt).Scan(&ml.ID, &ml.Email, &ml.Token, &ml.Used, &ml.ExpiresAt, &ml.CreatedAt)
	if err != nil {
		log.Printf("Error creating magic link for email %s: %v", email, err)
		return nil, err
	}
	return &ml, nil
}

func (r *MagicLinkRepo) GetByToken(token string) (*MagicLink, error) {
	var ml MagicLink
	err := r.db.QueryRow(`
		SELECT id, email, token, used, expires_at, created_at
		FROM magic_links
		WHERE token = $1
	`, token).Scan(&ml.ID, &ml.Email, &ml.Token, &ml.Used, &ml.ExpiresAt, &ml.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &ml, nil
}

func (r *MagicLinkRepo) MarkAsUsed(id int) error {
	_, err := r.db.Exec(`
		UPDATE magic_links
		SET used = true
		WHERE id = $1
	`, id)
	return err
}

func (r *MagicLinkRepo) DeleteExpired() error {
	_, err := r.db.Exec(`
		DELETE FROM magic_links
		WHERE expires_at < CURRENT_TIMESTAMP
	`)
	return err
}
