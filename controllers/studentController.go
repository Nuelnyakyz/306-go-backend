package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"306Web/go-backend/models"
	"306Web/go-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"306Web/go-backend/utils" // Assuming a utility for token verification
)

// GetStudentProfile fetches a student's profile by their ID
func GetStudentProfile(c *gin.Context) {
	// Get the token from the Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
		return
	}

	// Verify the token (e.g., using a helper function)
	studentID, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Find student by ID
	studentCollection := database.GetCollection("student_portal", "students")
	var student models.Student
	err = studentCollection.FindOne(c, bson.M{"_id": studentID}).Decode(&student)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Respond with the student profile and their registered courses
	c.JSON(http.StatusOK, gin.H{
		"student": student,
	})
}
