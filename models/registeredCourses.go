package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// RegisteredCourse represents a student's registered courses
type RegisteredCourse struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty"`    // MongoDB unique ID
	StudentID primitive.ObjectID   `bson:"studentId"`        // ID of the student
	CourseIDs []primitive.ObjectID `bson:"courseIds"`        // IDs of registered courses
	Timestamp int64                `bson:"timestamp"`        // Registration time
}
