package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB is the database connection
var DB *sql.DB

// InitiateConnection will start the communication with the database
func InitiateConnection() *sql.DB {
	DB, err := sql.Open("sqlite3", "jrdd.db")

	if err != nil {
		panic(fmt.Errorf("Fatal error connection to db: %s", err))
	}

	fmt.Println("Connection stablished succesfully")

	return DB
}

func getConnection() *sql.DB {
	return DB
}
