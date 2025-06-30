package db

import (
	"database/sql"
	"log"

	"github.com/NoCapCbas/webStash/internal/db/seed"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDB struct {
	DB *sql.DB
}

func NewSqliteDB(connStr string) (*SqliteDB, error) {
	log.Printf("Connecting to SqliteDB (%s)", connStr)

	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = seed.CreateUsersTable(db); err != nil {
		return nil, err
	}

	if err = seed.CreateSessionsTable(db); err != nil {
		return nil, err
	}

	if err = seed.CreateBookmarksTable(db); err != nil {
		return nil, err
	}

	if err = seed.CreateMagicLinksTable(db); err != nil {
		return nil, err
	}

	return &SqliteDB{DB: db}, nil
}
