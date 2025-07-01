package seed

import "database/sql"

func CreateBookmarksTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS bookmarks (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            user_id INTEGER NOT NULL,
            url TEXT NOT NULL,
            title TEXT,
            description TEXT,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
            public BOOLEAN NOT NULL DEFAULT 0,
            click_count INTEGER NOT NULL DEFAULT 0,
            FOREIGN KEY(user_id) REFERENCES users(id)
        );
        
        CREATE INDEX IF NOT EXISTS bookmarks_user_id_idx ON bookmarks(user_id);
        CREATE INDEX IF NOT EXISTS bookmarks_created_at_idx ON bookmarks(created_at);
    `)
	return err
}
