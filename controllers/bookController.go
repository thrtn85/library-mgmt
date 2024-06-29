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

// UpdateBook handles PUT requests to update book information
func UpdateBook(c *gin.Context) {
	var book models.Book
	bookID := c.Param("id")
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingBook models.Book
	if err := config.DB.First(&existingBook, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Update the book's information
	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	

	// Save the updated book to the database
	if err := config.DB.Save(&existingBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, existingBook)
}

// DeleteBook handles DELETE requests to delete a book
func DeleteBook(c *gin.Context) {
	var book models.Book
	bookID := c.Param("id")

	// Find the book by ID
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Delete the book from the database
	config.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}