package db

import (
	"database/sql"
	"log"

	"github.com/NoCapCbas/webStash/internal/db/seed"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

type User struct {
	Email          string
	MembershipType int
}

func NewPostgresDB(connStr string) (*PostgresDB, error) {
	log.Printf("Connecting to PostgresDB (%s)", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = seed.CreateUsersTable(db); err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

// CreateUser creates a new user with default membership type 0
func (p *PostgresDB) CreateUser(email string) error {
	_, err := p.db.Exec(`
        INSERT INTO users (email, membership_type) 
        VALUES ($1, $2)
        ON CONFLICT (email) DO NOTHING
    `, email, 0)
	return err
}

// GetUser retrieves a user by email
func (p *PostgresDB) GetUser(email string) (*User, error) {
	user := &User{}
	err := p.db.QueryRow(`
        SELECT email, membership_type 
        FROM users 
        WHERE email = $1
    `, email).Scan(&user.Email, &user.MembershipType)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}
