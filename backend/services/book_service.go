package services

import (
	"backend/models"
	"backend/repositories"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	CreateBook(input models.Book) (models.Book, error)
	GetBookByID(id string) (models.Book, error)
	UpdateBook(id string, input models.Book) (models.Book, error)
	DeleteBook(id string) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) CreateBook(input models.Book) (models.Book, error) {
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
	}
	err := s.repo.Create(&book)
	return book, err
}

func (s *bookService) GetBookByID(id string) (models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) UpdateBook(id string, input models.Book) (models.Book, error) {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}
	err = s.repo.Update(&book, models.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
	})
	return book, err
}
func (s *bookService) DeleteBook(id string) error {
	book, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(&book)
}
