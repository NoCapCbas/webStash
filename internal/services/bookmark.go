package services

import "github.com/NoCapCbas/webStash/internal/db/repos"

type BookmarkService struct {
	bookmarkRepo *repos.BookmarkRepo
}

func NewBookmarkService(bookmarkRepo *repos.BookmarkRepo) *BookmarkService {
	return &BookmarkService{bookmarkRepo: bookmarkRepo}
}

func (s *BookmarkService) CreateBookmark(bookmark *repos.Bookmark) error {
	return s.bookmarkRepo.Create(bookmark)
}

func (s *BookmarkService) UpdateBookmark(bookmark *repos.Bookmark) error {
	return s.bookmarkRepo.Update(bookmark)
}

func (s *BookmarkService) DeleteBookmark(id int, userID int) error {
	return s.bookmarkRepo.Delete(id, userID)
}

func (s *BookmarkService) GetBookmarkByID(id int) (*repos.Bookmark, error) {
	return s.bookmarkRepo.GetByID(id)
}

func (s *BookmarkService) GetBookmarksByUserID(userID int) ([]repos.Bookmark, error) {
	return s.bookmarkRepo.GetByUserID(userID)
}

func (s *BookmarkService) GetBookmarksByUserEmail(userEmail string) ([]repos.Bookmark, error) {
	return s.bookmarkRepo.GetByUserEmail(userEmail)
}

func (s *BookmarkService) IncrementClickCount(id int) error {
	return s.bookmarkRepo.IncrementClickCount(id)
}
