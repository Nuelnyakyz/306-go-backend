package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Student defines the schema for the students collection
type Student struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` // MongoDB unique ID
	Name     string             `bson:"name"`         // Student's full name
	Email    string             `bson:"email"`        // Email for authentication
	Password string             `bson:"password"`     // Hashed password
}
