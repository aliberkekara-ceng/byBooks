package repositories

import (
	"backend/config"
	"backend/models"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	Create(book *models.Book) error
	FindByID(id string) (models.Book, error)
	Update(book *models.Book, input models.Book) error
	Delete(book *models.Book) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := config.DB.Find(&books).Error
	return books, err
}

func (r *bookRepository) Create(book *models.Book) error {
	return config.DB.Create(book).Error
}

func (r *bookRepository) FindByID(id string) (models.Book, error) {
	var book models.Book
	err := config.DB.First(&book, id).Error
	return book, err
}

func (r *bookRepository) Update(book *models.Book, input models.Book) error {
	return config.DB.Model(book).Updates(input).Error
}

func (r *bookRepository) Delete(book *models.Book) error {
	return config.DB.Delete(book).Error
}
