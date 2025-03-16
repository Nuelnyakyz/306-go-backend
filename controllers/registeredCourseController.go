package controllers

import (
	"context"
	"net/http"
	"306Web/go-backend/database"
	"306Web/go-backend/models"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRegisteredCourses fetches all courses registered by the authenticated student
func GetRegisteredCourses(c *gin.Context) {
	// Retrieve userID from the context (set by AuthMiddleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	// Connect to the "registered_courses" collection
	registeredCoursesCollection := database.GetCollection("student_portal", "registered_courses")

	// Fetch registered courses for the given userID
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := registeredCoursesCollection.Find(ctx, bson.M{"student_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch registered courses"})
		return
	}
	defer cursor.Close(ctx)

	// Decode the results
	var registeredCourses []models.RegisteredCourse
	if err := cursor.All(ctx, &registeredCourses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse registered courses"})
		return
	}

	// Return the registered courses
	c.JSON(http.StatusOK, registeredCourses)
}
