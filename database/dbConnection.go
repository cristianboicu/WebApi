package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

//DB -> database variable
var DB *sql.DB

//Connect is the function for connecting to db
func Connect() error {
	var err error

	DB, err = sql.Open("sqlite3", "./sqlite-database.db")

	if err != nil {
		fmt.Println("Error connect to database")
	}

	CreateTaskTable()
	fmt.Println("Connection to db opened")
	return nil
}
