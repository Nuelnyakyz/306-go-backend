package routes

import (
	"github.com/gin-gonic/gin"
	"306Web/go-backend/controllers"
	"306Web/go-backend/utils"
)

// RegistrationRoutes defines routes for registering courses
func RegistrationRoutes(router *gin.Engine) {
	register := router.Group("/api/register")
	
	register.Use(utils.AuthMiddleware())
	//
	register.GET("/", controllers.GetRegisteredCourses)
	
}
