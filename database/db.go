package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is the MongoDB client instance
var DB *mongo.Client

// ConnectDB initializes the MongoDB connection
func ConnectDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB!")
	DB = client
}

// GetCollection gets a MongoDB collection
func GetCollection(dbName, colName string) *mongo.Collection {
	return DB.Database(dbName).Collection(colName)
}
