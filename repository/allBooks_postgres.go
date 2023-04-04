package repository

import (
	"fmt"
	"github.com/mereiamangeldin/bookStore/models"
	"gorm.io/gorm"
	"strings"
)

type allBooksPostgres struct {
	db *gorm.DB
}

func newAllBooksPostgres(db *gorm.DB) *allBooksPostgres {
	return &allBooksPostgres{db: db}
}

func (s *allBooksPostgres) GetAllBooks(sort string) ([]models.Book, error) {
	var books []models.Book
	err := s.db.Order(sort).Find(&books)
	if err.Error != nil {
		return nil, err.Error
	}
	return books, nil
}

func (s *allBooksPostgres) CreateBook(book models.Book) error {
	res := s.db.Create(&book)
	return res.Error
}

func (s *allBooksPostgres) GetBookByTitle(title string) (models.Book, error) {
	var book models.Book
	err := s.db.Where("LOWER(title) like ?", fmt.Sprintf("%s%%", strings.ToLower(title))).First(&book)
	if err.Error != nil {
		return models.Book{}, err.Error
	}
	return book, nil
}
