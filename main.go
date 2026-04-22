package main

import (
	"log"
	"net/http"
	"your_module/config" // Adjust the import path according to your module
)

func main() {
	// Initialize Redis connection
	redisClient, err := config.ConnectRedis()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// Check DB connection
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	// Error handling for route setup
	http.HandleFunc("/your-route", yourHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	log.Println("Server started on port 8080")
}