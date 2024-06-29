package routes

import (
	"github.com/lewislewin/testify-webserver/internal/controllers"
	"github.com/lewislewin/testify-webserver/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// TODO: Prune for the routes that are actually going to be needed
func SetupRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/auth/login", controllers.Login)
	r.POST("/auth/register", controllers.Register)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		// User routes
		protected.POST("/users", controllers.CreateUser)
		protected.GET("/users/:id", controllers.GetUser)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.DELETE("/users/:id", controllers.DeleteUser)
		protected.GET("/users", controllers.ListUsers)

		// Endpoint routes
		protected.POST("/endpoints", controllers.CreateEndpoint)
		protected.GET("/endpoints/:id", controllers.GetEndpoint)
		protected.PUT("/endpoints/:id", controllers.UpdateEndpoint)
		protected.DELETE("/endpoints/:id", controllers.DeleteEndpoint)
		protected.GET("/endpoints", controllers.ListEndpoints)

		// Company routes
		protected.GET("/companies/:id", controllers.GetCompany)
		protected.PUT("/companies/:id", controllers.UpdateCompany)
		protected.DELETE("/companies/:id", controllers.DeleteCompany)

		// TestSuite routes
		protected.POST("/testsuites", controllers.CreateTestSuite)
		protected.GET("/testsuites/:id", controllers.GetTestSuite)
		protected.PUT("/testsuites/:id", controllers.UpdateTestSuite)
		protected.DELETE("/testsuites/:id", controllers.DeleteTestSuite)
		protected.GET("/testsuites", controllers.ListTestSuites)
		protected.GET("/testsuites/:id/testresults", controllers.ListTestResults)
		protected.GET("/testsuites/:id/tests", controllers.ListTests)

		// Test routes TODO: this looks weird! I think it all needs to be /testsuites
		protected.POST("/testsuites/:id/tests", controllers.CreateTest)
		protected.GET("/tests/:id", controllers.GetTest)
		protected.PUT("/tests/:id", controllers.UpdateTest)
		protected.DELETE("/tests/:id", controllers.DeleteTest)

		// TestResult routes
		protected.GET("/testresults/:id", controllers.GetTestResult)

		// Other protected routes
		protected.GET("/me", controllers.GetCurrentUser)
	}
}
