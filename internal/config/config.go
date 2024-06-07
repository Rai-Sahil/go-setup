package config

import "os"

type Config struct {
	PostgresConnectionString string
}

func GetConfig() *Config {
	return &Config{
		PostgresConnectionString: os.Getenv("POSTGRES_CONNECTION_STRING"),
	}
}
