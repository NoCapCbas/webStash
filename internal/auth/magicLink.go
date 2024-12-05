package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

type MagicLink struct {
	Token     string
	Email     string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// TODO: create postgres cache for magic links
var tempCache = make(map[string]string)

func GenerateMagicLink(email string) (*MagicLink, error) {
	// Generate random token
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	rand.Read(b)
	for i := range b {
		b[i] = letterBytes[int(b[i])%len(letterBytes)]
	}
	token := base64.URLEncoding.EncodeToString(b)

	// Add token to cache
	tempCache[token] = email

	return &MagicLink{
		Token:     token,
		Email:     email,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}, nil
}

func ValidateMagicLink(token string) (string, error) {

	// TODO: Validate token
	email, ok := tempCache[token]
	if !ok {
		return "", errors.New("invalid or expired token")
	}

	// // check expiration
	// if time.Now().After(magicLink.ExpiresAt) {
	// 	return "", errors.New("invalid or expired token")
	// }

	return email, nil
}
