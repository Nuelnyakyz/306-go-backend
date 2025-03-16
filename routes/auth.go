package routes

import (
	"github.com/gin-gonic/gin"
	"306Web/go-backend/controllers" 
	
)

// AuthRoutes defines authentication-related routes
func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.Login) // Student login
		auth.POST("/register", controllers.Register) 
	}
}
