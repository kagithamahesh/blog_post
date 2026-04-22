package main

import (
	"BlogPost/config"
	"log"
)

func main() {
	// Initialize DB connection
	db := config.ConnectDB()
	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	// Initialize Redis connection
	redis := config.ConnectRedis()
	if redis == nil {
		log.Fatal("Failed to connect to Redis")
	}

}
