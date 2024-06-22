package routes

import (
	"testify-webserver/controllers"
	"testify-webserver/middlewears"

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

	// Protected routes
	protected := r.Group("/", middlewares.AuthMiddleware())
	{
		protected.GET("/me", controllers.GetCurrentUser)
	}
}
