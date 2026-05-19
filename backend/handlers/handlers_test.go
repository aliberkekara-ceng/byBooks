package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"backend/config"
	"backend/handlers"
	"backend/models"
	"backend/repositories"
	"backend/services"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter() *gin.Engine {
	// Initialize in-memory SQLite database for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to in-memory database: " + err.Error())
	}

	// Auto-migrate tables
	db.AutoMigrate(&models.Book{})
	config.DB = db

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// Instantiate layers
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	urlService := services.NewURLService()
	urlHandler := handlers.NewURLHandler(urlService)

	api := r.Group("/api")
	{
		api.GET("/books", bookHandler.FindBooks)
		api.POST("/books", bookHandler.CreateBook)
		api.GET("/books/:id", bookHandler.FindBook)
		api.PUT("/books/:id", bookHandler.UpdateBook)
		api.DELETE("/books/:id", bookHandler.DeleteBook)

		api.POST("/url-process", urlHandler.ProcessURL)
	}

	return r
}

func TestBookCRUD(t *testing.T) {
	router := setupTestRouter()

	var createdBook models.Book

	// 1. Test POST /api/books - Create Valid Book
	t.Run("Create Valid Book", func(t *testing.T) {
		payload := models.Book{
			Title:  "Test Driven Development",
			Author: "Kent Beck",
			Year:   2003,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusCreated {
			t.Errorf("Expected status 201, got %d. Response: %s", resp.Code, resp.Body.String())
		}

		err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
		if err != nil {
			t.Fatalf("Failed to unmarshal response: %v", err)
		}

		if createdBook.ID == 0 {
			t.Error("Expected assigned book ID, got 0")
		}
		if createdBook.Title != payload.Title {
			t.Errorf("Expected title %q, got %q", payload.Title, createdBook.Title)
		}
	})

	// 2. Test POST /api/books - Create Invalid Book (Missing Fields)
	t.Run("Create Invalid Book - Missing Title", func(t *testing.T) {
		payload := map[string]interface{}{
			"author": "Kent Beck",
			"year":   2003,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("POST", "/api/books", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusBadRequest {
			t.Errorf("Expected status 400, got %d", resp.Code)
		}
		if !strings.Contains(resp.Body.String(), "error") {
			t.Errorf("Expected validation error message, got %s", resp.Body.String())
		}
	})

	// 3. Test GET /api/books - Get All Books
	t.Run("Get All Books", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/books", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		var books []models.Book
		err := json.Unmarshal(resp.Body.Bytes(), &books)
		if err != nil {
			t.Fatalf("Failed to unmarshal books: %v", err)
		}

		if len(books) != 1 {
			t.Errorf("Expected 1 book, got %d", len(books))
		}
	})

	// 4. Test GET /api/books/:id - Get Single Existing Book
	t.Run("Get Book By ID - Success", func(t *testing.T) {
		url := "/api/books/" + strconv.Itoa(int(createdBook.ID))
		req, _ := http.NewRequest("GET", url, nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		var book models.Book
		json.Unmarshal(resp.Body.Bytes(), &book)
		if book.ID != createdBook.ID {
			t.Errorf("Expected book ID %d, got %d", createdBook.ID, book.ID)
		}
	})

	// 5. Test GET /api/books/:id - Get Non-existent Book
	t.Run("Get Book By ID - Not Found", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/books/9999", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", resp.Code)
		}
	})

	// 6. Test PUT /api/books/:id - Update Book Details
	t.Run("Update Book Details - Success", func(t *testing.T) {
		url := "/api/books/" + strconv.Itoa(int(createdBook.ID))
		payload := models.Book{
			Title:  "TDD By Example",
			Author: "Kent Beck",
			Year:   2003,
		}
		body, _ := json.Marshal(payload)

		req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		var updatedBook models.Book
		json.Unmarshal(resp.Body.Bytes(), &updatedBook)
		if updatedBook.Title != "TDD By Example" {
			t.Errorf("Expected updated title %q, got %q", "TDD By Example", updatedBook.Title)
		}
	})

	// 7. Test DELETE /api/books/:id - Delete Book Entry
	t.Run("Delete Book - Success", func(t *testing.T) {
		url := "/api/books/" + strconv.Itoa(int(createdBook.ID))
		req, _ := http.NewRequest("DELETE", url, nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", resp.Code)
		}

		// Verify deletion
		reqGet, _ := http.NewRequest("GET", url, nil)
		respGet := httptest.NewRecorder()
		router.ServeHTTP(respGet, reqGet)
		if respGet.Code != http.StatusNotFound {
			t.Errorf("Expected book to be deleted (404), got %d", respGet.Code)
		}
	})
}

func TestURLProcessorAPI(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name         string
		payload      models.URLRequest
		expectedStatus int
		expectedURL  string
		expectError  bool
	}{
		{
			name: "All Operation",
			payload: models.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://www.byfood.com/food-experiences",
			expectError:    false,
		},
		{
			name: "Canonical Operation",
			payload: models.URLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "canonical",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://BYFOOD.com/food-EXPeriences",
			expectError:    false,
		},
		{
			name: "Redirection Operation",
			payload: models.URLRequest{
				URL:       "https://example.com/some-path/",
				Operation: "redirection",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://www.byfood.com/some-path/",
			expectError:    false,
		},
		{
			name: "Invalid Operation Type",
			payload: models.URLRequest{
				URL:       "https://byfood.com",
				Operation: "invalid_operation",
			},
			expectedStatus: http.StatusBadRequest,
			expectedURL:    "",
			expectError:    true,
		},
		{
			name: "Empty URL",
			payload: models.URLRequest{
				URL:       "",
				Operation: "all",
			},
			expectedStatus: http.StatusBadRequest,
			expectedURL:    "",
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/url-process", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			if resp.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d. Response: %s", tt.expectedStatus, resp.Code, resp.Body.String())
			}

			if !tt.expectError {
				var response models.URLResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				if err != nil {
					t.Fatalf("Failed to decode response: %v", err)
				}
				if response.ProcessedURL != tt.expectedURL {
					t.Errorf("Expected URL %q, got %q", tt.expectedURL, response.ProcessedURL)
				}
			} else {
				if !strings.Contains(resp.Body.String(), "error") {
					t.Errorf("Expected error response, got: %s", resp.Body.String())
				}
			}
		})
	}
}
