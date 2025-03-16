package routes

import (
	"github.com/gin-gonic/gin"
	"306Web/go-backend/controllers"
)

// CourseRoutes defines routes related to courses
func CourseRoutes(router *gin.Engine) {
	courses := router.Group("/courses")
	{
		courses.GET("/", controllers.GetCourses) // Fetch all available courses
	}
}
