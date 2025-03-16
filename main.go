package main

import(
	"306Web/go-backend/database"
    "306Web/go-backend/routes"
)

func main() {
    // Initialize MongoDB connection
    database.ConnectDB()

    router := routes.SetupRouter()

	// Start the server
	router.Run(":8080")
}
