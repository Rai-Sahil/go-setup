package database

import (
	"database/sql"
	"github.com/Rai-Sahil/go-setup/internal/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	cfg := config.GetConfig()
	connectionString := cfg.PostgresConnectionString
	connectionString = "postgres://postgres:sql@localhost:5432/go-setup?sslmode=disable"
	return sql.Open("postgres", connectionString)
}
