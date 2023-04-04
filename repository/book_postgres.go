package repository

import (
	"github.com/mereiamangeldin/bookStore/models"
	"gorm.io/gorm"
)

type bookPostgres struct {
	db *gorm.DB
}

func newBookPostgres(db *gorm.DB) *bookPostgres {
	return &bookPostgres{db: db}
}

func (s *bookPostgres) DeleteBook(id int) error {
	var book models.Book
	result := s.db.First(&book, id)
	if result.Error != nil {
		return result.Error
	}
	result = s.db.Delete(&book)
	return result.Error
}

func (s *bookPostgres) UpdateBook(id int, bookRequest models.Book) error {
	var book models.Book
	result := s.db.Model(&book).Where("id = ?", id).Updates(bookRequest)
	return result.Error
}

func (s *bookPostgres) GetBookById(id int) (models.Book, error) {
	var book models.Book
	result := s.db.First(&book, id)
	if result.Error != nil {
		return models.Book{}, result.Error
	}
	return book, nil
}
