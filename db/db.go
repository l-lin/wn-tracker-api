package db

import (
	"log"
	"database/sql"
	"os"
)

type RowMapper interface {
	Scan(dest ...interface{}) error
}

func Connect() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("[x] Could not open the connection to the database. Reason: %s", err.Error())
	}
	return database
}
