package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"github.com/NoCapCbas/webStash/internal/db/repos"
)

type AuthService struct {
	magicLinkRepo *repos.MagicLinkRepo
	sessionRepo   *repos.SessionRepo
	userRepo      *repos.UserRepo
}

func NewAuthService(magicLinkRepo *repos.MagicLinkRepo, sessionRepo *repos.SessionRepo, userRepo *repos.UserRepo) *AuthService {
	return &AuthService{
		magicLinkRepo: magicLinkRepo,
		sessionRepo:   sessionRepo,
		userRepo:      userRepo,
	}
}

func (s *AuthService) GenerateMagicLink(email string) (*repos.MagicLink, error) {
	// Generate random token
	token, err := generateRandomToken()
	if err != nil {
		log.Println("Error generating token:", err)
		return nil, err
	}

	expiresAt := time.Now().Add(60 * time.Minute)
	ml, err := s.magicLinkRepo.Create(email, token, expiresAt)
	if err != nil {
		log.Println("Error creating magic link:", err)
		return nil, err
	}

	return &repos.MagicLink{
		Token:     ml.Token,
		Email:     ml.Email,
		CreatedAt: ml.CreatedAt,
		ExpiresAt: ml.ExpiresAt,
	}, nil
}

func (s *AuthService) ValidateMagicLink(token string) (string, error) {
	ml, err := s.magicLinkRepo.GetByToken(token)
	if err != nil {
		return "", errors.New("invalid token, please request a new one")
	}

	if time.Now().After(ml.ExpiresAt) {
		return "", errors.New("token expired, please request a new one")
	}

	s.magicLinkRepo.DeleteExpired()

	return ml.Email, nil
}

func (s *AuthService) CreateSession(userID int) (*repos.Session, error) {
	token, err := generateRandomToken()
	if err != nil {
		return nil, err
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	sess, err := s.sessionRepo.Create(userID, token, expiresAt)
	if err != nil {
		return nil, err
	}

	return &repos.Session{
		Token:     sess.Token,
		UserID:    sess.UserID,
		CreatedAt: sess.CreatedAt,
		ExpiresAt: sess.ExpiresAt,
	}, nil
}

func (s *AuthService) ValidateSession(token string) (string, error) {
	sess, err := s.sessionRepo.GetByToken(token)
	if err != nil {
		return "", errors.New("invalid session")
	}

	if time.Now().After(sess.ExpiresAt) {
		return "", errors.New("session expired, please login again")
	}

	user, err := s.userRepo.GetByID(sess.UserID)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}

func (s *AuthService) GenerateSessionToken(email string) (string, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	sess, err := s.CreateSession(user.ID)
	if err != nil {
		return "", err
	}
	return sess.Token, nil
}

func generateRandomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *AuthService) CreateUser(email string) error {
	// create user if not exists
	_, err := s.userRepo.Create(email, 0)
	return err
}
