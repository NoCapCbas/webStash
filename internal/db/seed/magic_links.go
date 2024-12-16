package seed

import (
	"database/sql"
	"log"
)

func CreateMagicLinksTable(db *sql.DB) error {
	log.Println("Creating magic links table")
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS magic_links (
            id SERIAL PRIMARY KEY,
            email VARCHAR(255) NOT NULL,
            token VARCHAR(255) UNIQUE NOT NULL,
            used_at TIMESTAMP WITH TIME ZONE,
            expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE INDEX IF NOT EXISTS magic_links_token_idx ON magic_links(token);
        CREATE INDEX IF NOT EXISTS magic_links_email_idx ON magic_links(email);
    `)
	if err != nil {
		log.Println("Error creating magic links table:", err)
		return err
	}
	log.Println("Magic links table created")
	return nil
}
