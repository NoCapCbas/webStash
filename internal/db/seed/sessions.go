package seed

import "database/sql"

func CreateSessionsTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS sessions (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            token TEXT UNIQUE NOT NULL,
            expires_at DATETIME NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY(user_id) REFERENCES users(id)
        );
        
        CREATE INDEX IF NOT EXISTS sessions_token_idx ON sessions(token);
        CREATE INDEX IF NOT EXISTS sessions_user_id_idx ON sessions(user_id);     `)
	return err
}
