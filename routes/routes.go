package routes

import (
	"testify-webserver/controllers"
	"testify-webserver/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.Register)

	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.GET("/users", controllers.ListUsers)

	// Endpoint routes
	r.POST("/endpoints", controllers.CreateEndpoint)
	r.GET("/endpoints/:id", controllers.GetEndpoint)
	r.PUT("/endpoints/:id", controllers.UpdateEndpoint)
	r.DELETE("/endpoints/:id", controllers.DeleteEndpoint)
	r.GET("/endpoints", controllers.ListEndpoints)

	// Protected routes
	protected := r.Group("/", middlewares.AuthMiddleware())
	{
		protected.GET("/me", controllers.GetCurrentUser)
	}
}
