package db

import (
	"database/sql"
	"log"

	"github.com/NoCapCbas/webStash/internal/db/seed"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(connStr string) (*PostgresDB, error) {
	log.Printf("Connecting to PostgresDB (%s)", connStr)

	db, err := sql.Open("postgres", connStr)
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

	return &PostgresDB{DB: db}, nil
}
