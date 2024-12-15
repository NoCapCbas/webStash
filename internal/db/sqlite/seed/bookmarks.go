package seed

import "database/sql"

func CreateBookmarksTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS bookmarks (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
			url TEXT NOT NULL,
			title TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}
