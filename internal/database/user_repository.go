package database

import (
	"github.com/Rai-Sahil/go-setup/internal/models"
)

func CreateTable() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
    )`)
	if err != nil {
		return err
	}

	return nil
}

func InsertUser(user models.User) error {
	db, err := Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, user.Name, user.Email)
	if err != nil {
		return err
	}

	return nil
}
