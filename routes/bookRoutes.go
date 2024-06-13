// bookRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/bookstore/controllers"
)

// sests up the routes for book-related endpoints
func BookRoutes(router *gin.Engine) {
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
}