package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var connection *sql.DB

func GetConnection() *sql.DB {

	if connection == nil {
		connection = createConnection()
	}

	fmt.Println("Database connected...")
	return connection
}

func createConnection() *sql.DB {

	connection, err := sql.Open("sqlite3", "data.sqlite")

	if err != nil {
		panic(err)
	}

	return connection
}
