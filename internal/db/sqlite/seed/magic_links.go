package seed

import "database/sql"

func CreateMagicLinksTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS magic_links (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
			email TEXT NOT NULL,
			token TEXT NOT NULL UNIQUE,
			used BOOLEAN DEFAULT FALSE,
			expires_at DATETIME NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}
