package service

import (
	"github.com/mereiamangeldin/bookStore/models"
	"github.com/mereiamangeldin/bookStore/repository"
)

type IBook interface {
	DeleteBook(id int) error
	UpdateBook(id int, book models.Book) error
	GetBookById(id int) (models.Book, error)
}
type IAllBooks interface {
	GetAllBooks(sort string) ([]models.Book, error)
	CreateBook(book models.Book) error
	GetBookByTitle(title string) (models.Book, error)
}

type Service struct {
	IBook
	IAllBooks
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		IAllBooks: newAllBooksService(repo.IAllBooks),
		IBook:     newBookService(repo.IBook),
	}
}
