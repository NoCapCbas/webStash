package seed

import (
	"database/sql"
	"log"
)

func CreateSessionsTable(db *sql.DB) error {
	log.Println("Creating sessions table")
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS sessions (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL REFERENCES users(id),
            token VARCHAR(255) UNIQUE NOT NULL,
            expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE INDEX IF NOT EXISTS sessions_token_idx ON sessions(token);
        CREATE INDEX IF NOT EXISTS sessions_user_id_idx ON sessions(user_id);
    `)
	if err != nil {
		log.Println("Error creating sessions table:", err)
		return err
	}
	log.Println("Sessions table created")
	return nil
}
