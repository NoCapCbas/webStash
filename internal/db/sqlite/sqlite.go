package sqlite

import (
	"database/sql"

	"github.com/NoCapCbas/webStash/internal/db/sqlite/seed"
	_ "github.com/mattn/go-sqlite3"
)

// DB represents the database connection
type DB struct {
	DB *sql.DB
}

// New creates a new database connection
func New(dbPath string) (*DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// InitSchema initializes the database schema
func (db *DB) InitSchema() error {
	if err := seed.CreateUsersTable(db); err != nil {
		return err
	}
	if err := seed.CreateSessionsTable(db); err != nil {
		return err
	}
	if err := seed.CreateBookmarksTable(db); err != nil {
		return err
	}
	if err := seed.CreateMagicLinksTable(db); err != nil {
		return err
	}
	return nil
}

// Transaction executes a function within a database transaction
func (db *DB) Transaction(fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
