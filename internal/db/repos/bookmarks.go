package repos

import (
	"database/sql"
	"time"
)

type Bookmark struct {
	ID          int
	UserID      int
	URL         string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Public      bool
	ClickCount  int
}

type BookmarkRepo struct {
	db *sql.DB
}

func NewBookmarkRepo(db *sql.DB) *BookmarkRepo {
	return &BookmarkRepo{db: db}
}

func (r *BookmarkRepo) Create(bookmark *Bookmark) error {
	return r.db.QueryRow(`
		INSERT INTO bookmarks (user_id, url, title, description, public)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at, click_count
	`, bookmark.UserID, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public).
		Scan(&bookmark.ID, &bookmark.CreatedAt, &bookmark.UpdatedAt, &bookmark.ClickCount)
}

func (r *BookmarkRepo) GetByID(id int) (*Bookmark, error) {
	var b Bookmark
	err := r.db.QueryRow(`
		SELECT id, user_id, url, title, description, created_at, updated_at, public, click_count
		FROM bookmarks
		WHERE id = $1
	`, id).Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.Description, &b.CreatedAt, &b.UpdatedAt, &b.Public, &b.ClickCount)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookmarkRepo) GetByUserID(userID int) ([]Bookmark, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, url, title, description, created_at, updated_at, public, click_count
		FROM bookmarks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var b Bookmark
		if err := rows.Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.Description, &b.CreatedAt, &b.UpdatedAt, &b.Public, &b.ClickCount); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, nil
}

func (r *BookmarkRepo) Update(bookmark *Bookmark) error {
	_, err := r.db.Exec(`
		UPDATE bookmarks
		SET url = $1, title = $2, description = $3, public = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5 AND user_id = $6
	`, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, bookmark.ID, bookmark.UserID)
	return err
}

func (r *BookmarkRepo) IncrementClickCount(id int) error {
	_, err := r.db.Exec(`
		UPDATE bookmarks
		SET click_count = click_count + 1
		WHERE id = $1
	`, id)
	return err
}

func (r *BookmarkRepo) Delete(id, userID int) error {
	_, err := r.db.Exec(`
		DELETE FROM bookmarks
		WHERE id = $1 AND user_id = $2
	`, id, userID)
	return err
}
