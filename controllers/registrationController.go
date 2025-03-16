package controllers

import (
	"net/http"
	"306Web/go-backend/models"
	"306Web/go-backend/utils"
	"306Web/go-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

// Register registers a new student
func Register(c *gin.Context) {
	var newStudent struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// Bind the input JSON data to the newStudent struct
	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Check if student already exists by email
	studentCollection := database.GetCollection("student_portal", "students")
	var existingStudent models.Student
	err := studentCollection.FindOne(c, bson.M{"email": newStudent.Email}).Decode(&existingStudent)
	if err != nil && err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// If a student with this email already exists, return an error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already in use"})
		return
	}

	// Create a new student model
	student := models.Student{
		ID:       primitive.NewObjectID(),
		Name:     newStudent.Name,
		Email:    newStudent.Email,
		Password: newStudent.Password, // Consider hashing the password here
	}

	// Insert the new student into the database
	_, err = studentCollection.InsertOne(c, student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register student"})
		return
	}

	// Generate a token for the new student
	token, err := utils.GenerateToken(student.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return a success message and the generated token
	c.JSON(http.StatusOK, gin.H{
		"message": "Student registered successfully",
		"token":   token,
	})
}
