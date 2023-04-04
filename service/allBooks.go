package service

import (
	"github.com/mereiamangeldin/bookStore/models"
	"github.com/mereiamangeldin/bookStore/repository"
)

type allBooksService struct {
	repo repository.IAllBooks
}

func newAllBooksService(repo repository.IAllBooks) *allBooksService {
	return &allBooksService{repo: repo}
}

func (s *allBooksService) GetAllBooks(sort string) ([]models.Book, error) {
	return s.repo.GetAllBooks(sort)
}

func (s *allBooksService) CreateBook(book models.Book) error {
	return s.repo.CreateBook(book)
}

func (s *allBooksService) GetBookByTitle(title string) (models.Book, error) {
	return s.repo.GetBookByTitle(title)
}
