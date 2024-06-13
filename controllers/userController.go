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

// UpdateUser handles PUT requests to update user information
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := config.DB.First(&existingUser, user.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update the user's information
	existingUser.Name = user.Name
	existingUser.Email = user.Email

	// Save the updated user to the database
	config.DB.Save(&existingUser)

	c.JSON(http.StatusOK, existingUser)
}

// DeleteUser handles DELETE requests to delete a user
func DeleteUser(c *gin.Context) {
	var user models.User
	userID := c.Param("id")

	// Find the user by ID
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Delete the user from the database
	config.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}