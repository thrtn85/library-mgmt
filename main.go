package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/bookstore/config"
	"github.com/thrtn85/bookstore/routes"
)

func main() {
	router := gin.Default()

	// Connect to the database and automatically migrate the schema
	config.ConnectDatabase()

	// Register book routes
	routes.BookRoutes(router)
	routes.UserRoutes(router)

	// Start the server on port 8080
	router.Run(":8080")
}
