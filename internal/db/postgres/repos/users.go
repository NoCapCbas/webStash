package repos

import (
	"database/sql"

	"github.com/NoCapCbas/webStash/internal/db/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Create inserts a new user and returns the created user with ID
func (r *UserRepo) Create(email string, membershipType int) (*models.User, error) {
	var user models.User

	// First try to get existing user
	err := r.db.QueryRow(`
		SELECT id, email, membership_type, created_at 
		FROM users 
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)

	if err == nil {
		// User already exists, return it
		return &user, nil
	}

	// If user doesn't exist, create new one
	err = r.db.QueryRow(`
		INSERT INTO users (email, membership_type)
		VALUES ($1, $2)
		RETURNING id, email, membership_type, created_at
	`, email, membershipType).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByID retrieves a user by their ID
func (r *UserRepo) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, email, membership_type, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by their email
func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, email, membership_type, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update modifies an existing user's information
func (r *UserRepo) Update(user *models.User) error {
	_, err := r.db.Exec(`
		UPDATE users
		SET email = $1, membership_type = $2
		WHERE id = $3
	`, user.Email, user.MembershipType, user.ID)
	return err
}

// Delete removes a user by their ID
func (r *UserRepo) Delete(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM users
		WHERE id = $1
	`, id)
	return err
}
