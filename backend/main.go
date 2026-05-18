package main

import (
	"backend/config"
	_ "backend/docs"
	"backend/handlers"
	"backend/repositories"
	"backend/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Printf("[API Request] %s - %s | Status: %d | Duration: %v | IP: %s",
			reqMethod, reqUri, statusCode, latencyTime, clientIP,
		)
	}
}

// @title ByFood Assignment API
// @version 1.0
// @description Interactive API documentation for Book CRUD and URL cleanup service.
// @host localhost:8080
// @BasePath /api
func main() {
	config.ConnectDatabase()

	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	urlService := services.NewURLService()
	urlHandler := handlers.NewURLHandler(urlService)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())
	r.Use(LoggerMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/books", bookHandler.FindBooks)
		api.POST("/books", bookHandler.CreateBook)
		api.GET("/books/:id", bookHandler.FindBook)
		api.PUT("/books/:id", bookHandler.UpdateBook)
		api.DELETE("/books/:id", bookHandler.DeleteBook)

		api.POST("/url-process", urlHandler.ProcessURL)
	}

	log.Println("Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
