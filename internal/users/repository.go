// repository containing data access logic
package users

import (
	"database/sql"
	"fmt"
)

type UserPostgresRepository interface {
	Create(*User) (*User, error)
	Update(user *User) error
	GetByID(id int) (*User, error)
	GetByEmail(email string) (*User, error)
	Verify(user *User) error
}

type UserPostgresRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserPostgresRepository {
	return &UserPostgresRepositoryImpl{DB: db}
}

func (r *UserPostgresRepositoryImpl) Create(user *User) (*User, error) {
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

func (r *UserPostgresRepositoryImpl) Verify(user *User) error {
	fmt.Println("Verified user")
	return nil
}
