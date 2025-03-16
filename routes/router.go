package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes all routes for the application
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow cookies or authentication headers
	}))

	// Add route groups
	AuthRoutes(router)        // Authentication routes
	StudentRoutes(router)     // Student routes
	CourseRoutes(router)      // Course routes
	RegistrationRoutes(router) // Registration routes

	return router
}
