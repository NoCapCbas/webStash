package repos

import (
	"database/sql"
	"log"
	"time"
)

type Bookmark struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tags        []Tag     `json:"tags"`
	Public      bool      `json:"public"`
	ClickCount  int       `json:"click_count"`
}

type BookmarkRepo struct {
	db      *sql.DB
	tagRepo *TagRepo
}

func NewBookmarkRepo(db *sql.DB) *BookmarkRepo {
	return &BookmarkRepo{db: db, tagRepo: NewTagRepo(db)}
}

func (r *BookmarkRepo) Create(bookmark *Bookmark) error {
	log.Println("Creating bookmark: ", bookmark)
	_, err := r.db.Exec(`
		INSERT INTO bookmarks (user_id, url, title, description, public)
		VALUES ($1, $2, $3, $4, $5)
	`, bookmark.UserID, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, bookmark.Tags)
	if err != nil {
		log.Printf("Error creating bookmark(%d, %s, %s, %s, %t): %v", bookmark.UserID, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, err)
	}
	return err
}

func (r *BookmarkRepo) GetByID(id int) (*Bookmark, error) {
	var b Bookmark
	err := r.db.QueryRow(`
		SELECT id, user_id, url, title, description, created_at, updated_at, public, click_count, tags
		FROM bookmarks
		WHERE id = $1
	`, id).Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.Description, &b.CreatedAt, &b.UpdatedAt, &b.Public, &b.ClickCount, &b.Tags)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *BookmarkRepo) GetByUserID(userID int) ([]Bookmark, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, url, title, description, created_at, updated_at, public, click_count, tags
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
		if err := rows.Scan(&b.ID, &b.UserID, &b.URL, &b.Title, &b.Description, &b.CreatedAt, &b.UpdatedAt, &b.Public, &b.ClickCount, &b.Tags); err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, b)
	}
	return bookmarks, nil
}

func (r *BookmarkRepo) GetByUserEmail(userEmail string) ([]Bookmark, error) {
	var userID int
	err := r.db.QueryRow(`
		SELECT id FROM users WHERE email = $1
	`, userEmail).Scan(&userID)
	if err != nil {
		return nil, err
	}
	return r.GetByUserID(userID)
}

func (r *BookmarkRepo) Update(bookmark *Bookmark) error {
	_, err := r.db.Exec(`
		UPDATE bookmarks
		SET url = $1, title = $2, description = $3, public = $4, tags = $5, updated_at = CURRENT_TIMESTAMP
		WHERE id = $6
	`, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, bookmark.Tags, bookmark.ID)
	if err != nil {
		log.Printf("Error updating bookmark(%d, %s, %s, %s, %t): %v", bookmark.ID, bookmark.URL, bookmark.Title, bookmark.Description, bookmark.Public, err)
	}
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
	if err != nil {
		log.Printf("Error deleting bookmark(%d, %d): %v", id, userID, err)
	}

	// delete tags
	err = r.tagRepo.DeleteAllByBookmarkID(id)
	if err != nil {
		log.Printf("Error deleting tags for bookmark(%d): %v", id, err)
	}

	return err
}
