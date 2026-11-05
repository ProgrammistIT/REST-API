package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {
	connectionString := os.Getenv("DATABASE_URL")

	if connectionString == "" {
		connectionString = "postgres://mac@localhost:5432/threat_db?sslmode=disable"
	}

	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Successful connection to PostgreSQL")
	return db, nil
}
