package env

import (
	"log"

	"github.com/joho/godotenv"
)

// GetValue is function to get .env by name
func GetValue(key string) string {
	var env map[string]string
	env, err := godotenv.Read()

	if err != nil {
		log.Fatal("env: Error loading .env file")
	}

	envValue, ok := env[key]

	if !ok {
		log.Fatal("env: " + key + " is not existed")
	}

	return envValue
}
