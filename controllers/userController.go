// userController.go
package controllers

import (
	"net/http"

	"github.com/thrtn85/bookstore/config"
	"github.com/thrtn85/bookstore/models"

	"github.com/gin-gonic/gin"
)

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {
	var user models.User
	// Bind the JSON payload to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the new user to the database
	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)
	}

	// GetUsers handles GET requests to retrieve all users
	func GetUsers(c *gin.Context) {
	var users []models.User
	// Fetch all users from the database
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
