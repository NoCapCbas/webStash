package repos

import (
	"database/sql"

	"github.com/NoCapCbas/webStash/internal/db/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) Create(user *models.User) error {
	// first check if user exists
	_, err := r.GetByEmail(user.Email)
	if err == nil {
		return nil
	}

	// if user doesn't exist, create new one

	_, err = r.db.Exec(`
		INSERT INTO users (email, membership_type)
		VALUES ($1, $2)
	`, user.Email, user.MembershipType)

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, email, membership_type, created_at FROM users WHERE id = $1
	`, id).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)
	return &user, err
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRow(`
		SELECT id, email, membership_type, created_at FROM users WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.MembershipType, &user.CreatedAt)
	return &user, err
}

func (r *UserRepo) Update(user *models.User) error {
	_, err := r.db.Exec(`
		UPDATE users SET email = $1, membership_type = $2 WHERE id = $3
	`, user.Email, user.MembershipType, user.ID)
	return err
}
