// bookController.go
package controllers

import (
	"net/http"

	"github.com/thrtn85/bookstore/config"
	"github.com/thrtn85/bookstore/models"

	"github.com/gin-gonic/gin"
)

// CreateBook handles POST requests to create a new book
func CreateBook(c *gin.Context) {
	var book models.Book
	// Bind the JSON payload to the book struct
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the new book to the database
	config.DB.Create(&book)
	c.JSON(http.StatusOK, book)
	}

	// GetBooks handles GET requests to retrieve all books
	func GetBooks(c *gin.Context) {
	var books []models.Book
	// Fetch all books from the database
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}
