package database

import (
	"database/sql"
	"github.com/Rai-Sahil/go-setup/internal/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	cfg := config.GetConfig()
	connectionString := cfg.PostgresConnectionString
	return sql.Open("postgres", connectionString)
}
