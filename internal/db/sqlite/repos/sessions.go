package repos

import (
	"database/sql"
	"time"
)

type Session struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type SessionRepo struct {
	db *sql.DB
}

func NewSessionRepo(db *sql.DB) *SessionRepo {
	return &SessionRepo{db: db}
}

func (r *SessionRepo) Create(userID int, token string, expiresAt time.Time) (*Session, error) {
	var session Session
	err := r.db.QueryRow(`
		INSERT INTO sessions (user_id, token, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, token, expires_at, created_at
	`, userID, token, expiresAt).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepo) GetByToken(token string) (*Session, error) {
	var session Session
	err := r.db.QueryRow(`
		SELECT id, user_id, token, expires_at, created_at
		FROM sessions
		WHERE token = $1 AND expires_at > CURRENT_TIMESTAMP
	`, token).Scan(&session.ID, &session.UserID, &session.Token, &session.ExpiresAt, &session.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepo) DeleteByToken(token string) error {
	_, err := r.db.Exec(`
		DELETE FROM sessions
		WHERE token = $1
	`, token)
	return err
}

func (r *SessionRepo) DeleteByUserID(userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM sessions
		WHERE user_id = $1
	`, userID)
	return err
}

func (r *SessionRepo) DeleteExpired() error {
	_, err := r.db.Exec(`
		DELETE FROM sessions
		WHERE expires_at < CURRENT_TIMESTAMP
	`)
	return err
}
