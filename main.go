package main

import (
	"github.com/joho/godotenv"
	"github.com/okiprianto23/login-oauth/config"
	"github.com/okiprianto23/login-oauth/routes"
	"log"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	config.ConnectDB()

	// Run migrations
	config.MigrateDB()

	// Setup routes
	r := routes.SetupRouter()

	// Start server
	r.Run(":9899")
}
