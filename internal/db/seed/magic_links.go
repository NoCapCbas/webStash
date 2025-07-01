package seed

import "database/sql"

func CreateMagicLinksTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS magic_links (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL,
            token TEXT UNIQUE NOT NULL,
            used BOOLEAN DEFAULT 0,
            expires_at DATETIME NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
        
        CREATE INDEX IF NOT EXISTS magic_links_token_idx ON magic_links(token);
        CREATE INDEX IF NOT EXISTS magic_links_email_idx ON magic_links(email);     `)
	return err
}
