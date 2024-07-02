// bookRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/library-mgmt/controllers"
)

// sests up the routes for book-related endpoints
func UserRoutes(router *gin.Engine) {
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}