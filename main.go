package main

import (
	"testify-webserver/database"
	"testify-webserver/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the database
	database.InitDB()

	// Set up routes
	routes.SetupRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
