package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresConnectionString string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PostgresConnectionString: os.Getenv("POSTGRESQL_CONNECTION_STRING"),
	}
}
