package controllers

import (
	"net/http"
	"306Web/go-backend/models"
	"306Web/go-backend/utils"
	"306Web/go-backend/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login authenticates a user and returns a JWT token
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Query the student from the database
	studentCollection := database.GetCollection("student_portal", "students")
	var student models.Student
	err := studentCollection.FindOne(c, bson.M{"email": loginData.Email, "password": loginData.Password}).Decode(&student)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(student.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
