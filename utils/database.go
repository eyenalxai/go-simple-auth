package utils

import (
	"database/sql"
	"errors"
	users "go-simple-auth/postgresql"
	"os"
)

func GetDbConnection() (*users.Queries, error) {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return nil, errors.New("environment variable DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	queries := users.New(db)

	return queries, nil
}
