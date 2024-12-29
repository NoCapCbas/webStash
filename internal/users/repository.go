// repository containing data access logic
package users

import (
	"database/sql"
	"github.com/NoCapCbas/webStash/internal/users/model"
)

UserPostgresRepository interface {
	Create(*User) (*User, error)
	Update(user *User) error
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
}

type UserPostgresRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserPostgresRepositoryImpl) Create() (*User, error) {
	fmt.Println("Created user")
	return nil, nil
}

func (r *UserPostgresRepositoryImpl) Update(user *User) error {
	fmt.Println("Updated user")
	return nil
}

func (r *UserPostgresRepositoryImpl) GetByID(id int) (*User, error) {
	fmt.Println("Get user by id")
	return nil, nil
}

func (r *UserPostgresRepositoryImpl) GetByEmail(email string) (*User, error) {
	fmt.Println("Get user by email")
	return nil, nil
}
