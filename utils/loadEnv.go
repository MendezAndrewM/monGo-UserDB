package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Opens .env file for envVar access
func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}
