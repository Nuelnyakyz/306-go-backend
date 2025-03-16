package routes

import (
	"github.com/gin-gonic/gin"
	"306Web/go-backend/controllers"
	"306Web/go-backend/utils"
)

// StudentRoutes defines routes related to student profiles
func StudentRoutes(router *gin.Engine) {
	students := router.Group("/students")
	{
		students.Use(utils.AuthMiddleware())
		
		students.GET("/:id", controllers.GetStudentProfile) // Fetch student profile (including registered courses)
	}
}
