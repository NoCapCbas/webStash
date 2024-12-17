package repos

import (
	"database/sql"
	"time"
)

type Tag struct {
	ID         int       `json:"id"`
	BookmarkID int       `json:"bookmark_id"`
	Tag        string    `json:"tag"`
	CreatedAt  time.Time `json:"created_at"`
}

type TagRepo struct {
	db *sql.DB
}

func NewTagRepo(db *sql.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) Create(tag *Tag) error {
	_, err := r.db.Exec(`
		INSERT INTO tags (bookmark_id, tag)
		VALUES ($1, $2)
	`, tag.BookmarkID, tag.Tag)
	return err
}

func (r *TagRepo) GetAllByBookmarkID(bookmarkID int) ([]Tag, error) {
	rows, err := r.db.Query(`
		SELECT id, bookmark_id, tag, created_at
		FROM tags
		WHERE bookmark_id = $1
	`, bookmarkID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.BookmarkID, &tag.Tag, &tag.CreatedAt)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *TagRepo) DeleteAllByBookmarkID(bookmarkID int) error {
	_, err := r.db.Exec(`
		DELETE FROM tags
		WHERE bookmark_id = $1
	`, bookmarkID)
	return err
}

func (r *TagRepo) Delete(id int) error {
	_, err := r.db.Exec(`
		DELETE FROM tags
		WHERE id = $1
	`, id)
	return err
}
