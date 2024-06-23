package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LogMiddleware logs the details of each request and response
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Log details
		duration := time.Since(startTime)
		log.Printf("MIDDLEWEAR EXAMPLE FOR LOGGING %s %s %d %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
