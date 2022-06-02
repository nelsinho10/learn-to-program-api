package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnv returns the value of the environment variable named by the key.
// If the variable is not present, GetEnv returns the default value.
func GetEnv(key, defaultValue string) string {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}

	return defaultValue
}
