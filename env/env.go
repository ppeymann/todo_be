package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// GetStringDefault returns string from environment variable for specific key or returns expected default value.
func GetStringDefault(key string, def string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	val := os.Getenv(string(key))
	if len(val) == 0 {
		return def
	}

	return val
}

func IsProduction() bool {
	val := GetStringDefault("GIN_MODE", "debug")
	if val == "release" {
		return true
	}
	return false
}
