// bookRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/library-mgmt/controllers"
)

// sests up the routes for book-related endpoints
func BookRoutes(router *gin.Engine) {
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)
}
