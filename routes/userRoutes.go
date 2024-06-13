// bookRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/bookstore/controllers"
)

// sests up the routes for book-related endpoints
func UserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
}