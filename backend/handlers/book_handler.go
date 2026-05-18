package handlers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service services.BookService
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// FindBooks godoc
// @Summary Get all books
// @Description Retrieve a list of all books in the library
// @Tags books
// @Produce json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (h *BookHandler) FindBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Add a new book
// @Description Create a new book entry in the library
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book Data"
// @Success 201 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.CreateBook(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// FindBook godoc
// @Summary Get a book by ID
// @Description Retrieve details of a single book by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} map[string]string
// @Router /books/{id} [get]
func (h *BookHandler) FindBook(c *gin.Context) {
	id := c.Param("id")
	book, err := h.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update an existing book's details by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Updated Book Data"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.service.UpdateBook(id, input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book entry from the library by its ID
// @Tags books
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
