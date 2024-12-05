package auth

import (
	"crypto/rand"
	"errors"
	"time"
)

type Session struct {
	Token     string
	Email     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// TODO: Store sessions in database
var sessions = make(map[string]Session)

func GenerateSessionToken(email string) (string, error) {
	// Generate random token
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	rand.Read(b)
	for i := range b {
		b[i] = letterBytes[int(b[i])%len(letterBytes)]
	}
	token := string(b)

	// Store session in database
	sessions[token] = Session{
		Token:     token,
		Email:     email,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24),
	}

	return token, nil
}

func ValidateSession(token string) (string, error) {
	session, ok := sessions[token]
	if !ok {
		return "", errors.New("invalid session token")
	}
	return session.Email, nil
}
