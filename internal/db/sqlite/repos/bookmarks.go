package repos

import (
	"database/sql"

	"github.com/NoCapCbas/webStash/internal/db/models"
)

type BookmarkRepo struct {
	db *sql.DB
}

func NewBookmarkRepo(db *sql.DB) *BookmarkRepo {
	return &BookmarkRepo{db}
}

func (r *BookmarkRepo) Create(bookmark *models.Bookmark) error {
	_, err := r.db.Exec(`
		INSERT INTO bookmarks (user_id, url, title, description, public)
		VALUES ($1, $2, $3, $4, $5)
	`, bookmark.UserID, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public)
	return err
}

func (r *BookmarkRepo) GetByID(id int) (*models.Bookmark, error) {
	var bookmark models.Bookmark
	err := r.db.QueryRow(`
		SELECT id, user_id, url, title, description, created_at, updated_at, public, click_count
		FROM bookmarks
		WHERE id = $1
	`, id).Scan(&bookmark.ID, &bookmark.UserID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.CreatedAt, &bookmark.UpdatedAt, &bookmark.Public, &bookmark.ClickCount)
	return &bookmark, err
}

func (r *BookmarkRepo) GetByUserID(userID int) ([]models.Bookmark, error) {
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

	var bookmarks []models.Bookmark
	for rows.Next() {
		var bookmark models.Bookmark
		if err := rows.Scan(&bookmark.ID, &bookmark.UserID, &bookmark.URL, &bookmark.Title, &bookmark.Description, &bookmark.CreatedAt, &bookmark.UpdatedAt, &bookmark.Public, &bookmark.ClickCount); err != nil {
			return nil, err
		}
	}
	return bookmarks, nil
}

func (r *BookmarkRepo) GetByUserEmail(userEmail string) ([]models.Bookmark, error) {
	var userID int
	err := r.db.QueryRow(`
		SELECT id FROM users WHERE email = $1
	`, userEmail).Scan(&userID)
	if err != nil {
		return nil, err
	}
	return r.GetByUserID(userID)
}

func (r *BookmarkRepo) Update(bookmark *models.Bookmark) error {
	_, err := r.db.Exec(`
		UPDATE bookmarks SET url = $1, title = $2, description = $3, public = $4, click_count = $5 WHERE id = $6
	`, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, bookmark.ClickCount, bookmark.ID)
	return err
}

func (r *BookmarkRepo) IncrementClickCount(id int) error {
	_, err := r.db.Exec(`
		UPDATE bookmarks SET click_count = click_count + 1 WHERE id = $1
	`, id)
	return err
}

func (r *BookmarkRepo) Delete(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM bookmarks WHERE id = $1
	`, id)
	return err
}
