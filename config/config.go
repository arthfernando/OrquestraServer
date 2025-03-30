package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envVars = []string{
	"RABBITMQ_URL",
}

var envVar = map[string]string{}

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	for _, key := range envVars {
		value := os.Getenv(key)
		if value == "" {
			log.Fatalf("Missing env variable: %s", key)
		}
		envVar[key] = value
	}
}

func Get(name string) string {
	value := envVar[name]
	if value == "" {
		log.Fatalf("Wrong env variable: %s", name)
	}
	return value
}
