package service

import (
	"github.com/mereiamangeldin/bookStore/models"
	"github.com/mereiamangeldin/bookStore/repository"
)

type bookService struct {
	repo repository.IBook
}

func newBookService(repo repository.IBook) *bookService {
	return &bookService{repo: repo}
}

func (s *bookService) DeleteBook(id int) error {
	return s.repo.DeleteBook(id)
}

func (s *bookService) UpdateBook(id int, book models.Book) error {
	return s.repo.UpdateBook(id, book)
}

func (s *bookService) GetBookById(id int) (models.Book, error) {
	return s.repo.GetBookById(id)
}
