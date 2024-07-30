package main

import (
	"fmt"

	"github.com/j0n4t45d3v/crud/database"
)

func main() {
	con := database.GetConnection()

	queryTable := "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, username VARCHAR(20), password VARCHAR(10), email VARCHAR(50) )"
	_, err := con.Exec(queryTable)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table created")

}
