package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"306Web/go-backend/models"
	"306Web/go-backend/database"
	"go.mongodb.org/mongo-driver/bson"
)

// GetCourses fetches all available courses
func GetCourses(c *gin.Context) {
	courseCollection := database.GetCollection("student_portal", "courses")
	cursor, err := courseCollection.Find(c, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch courses"})
		return
	}
	defer cursor.Close(c)

	var courses []models.Course
	if err := cursor.All(c, &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to decode courses"})
		return
	}

	// Respond with the courses array directly
	c.JSON(http.StatusOK, courses)
}
