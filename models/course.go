package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Course defines the schema for the courses collection
type Course struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`  // MongoDB unique ID
	CourseCode string             `bson:"code"`     // Course code (e.g., CS101)
	CourseName string             `bson:"name"`     // Full course name
	Lecturer   string             `bson:"lecturer"`       // Name of the lecturer for the course
}
