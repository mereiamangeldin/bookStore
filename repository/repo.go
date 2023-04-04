package repository

import (
	"github.com/mereiamangeldin/bookStore/models"
	"gorm.io/gorm"
)

type IBook interface {
	DeleteBook(id int) error
	GetBookById(id int) (models.Book, error)
	UpdateBook(id int, book models.Book) error
}
type IAllBooks interface {
	GetAllBooks(sort string) ([]models.Book, error)
	CreateBook(book models.Book) error
	GetBookByTitle(title string) (models.Book, error)
}

type Repository struct {
	IBook
	IAllBooks
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IAllBooks: newAllBooksPostgres(db),
		IBook:     newBookPostgres(db)}
}
