package seed

import "database/sql"

func CreateBookmarksTable(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS bookmarks (
            id SERIAL PRIMARY KEY,
            user_id INTEGER NOT NULL REFERENCES users(id),
            url TEXT NOT NULL,
            title TEXT,
            description TEXT,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			public BOOLEAN NOT NULL DEFAULT FALSE,
			click_count INTEGER NOT NULL DEFAULT 0
        );
        
        CREATE INDEX IF NOT EXISTS bookmarks_user_id_idx ON bookmarks(user_id);
        CREATE INDEX IF NOT EXISTS bookmarks_created_at_idx ON bookmarks(created_at);
    `)
	return err
}
