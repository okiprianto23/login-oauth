package config

import (
	"fmt"
	"github.com/okiprianto23/login-oauth/models"
	"log"
)

// MigrateDB runs database migrations
func MigrateDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	fmt.Println("Database migration completed successfully!")
}
