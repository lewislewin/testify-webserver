package main

import (
	"github.com/lewislewin/testify-webserver/internal/middlewares"
	"github.com/lewislewin/testify-webserver/internal/routes"

	"github.com/lewislewin/testify-webserver/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Use logging middleware
	r.Use(middlewares.LogMiddleware())

	// Initialize the database
	database.InitDB()

	// Set up routes
	routes.SetupRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
