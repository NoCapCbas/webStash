package seed

import "database/sql"

func CreateUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			membership_type TEXT NOT NULL DEFAULT 'basic' CHECK(membership_type IN ('basic', 'pro')),
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`)
	return err
}
