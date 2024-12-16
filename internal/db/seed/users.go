package seed

import (
	"database/sql"
	"log"
)

func CreateUsersTable(db *sql.DB) error {
	log.Println("Creating users table")
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            email VARCHAR(255) UNIQUE NOT NULL,
            membership_type INTEGER NOT NULL DEFAULT 0,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE INDEX IF NOT EXISTS users_email_idx ON users(email);
    `)
	if err != nil {
		log.Println("Error creating users table:", err)
		return err
	}
	log.Println("Users table created")
	return nil
}
