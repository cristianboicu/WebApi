package database

import (
	"database/sql"
	"log"
)

//DB -> database variable
var DB *sql.DB

//Connect is the function for connecting to db
func Connect() error {
	var err error

	DB, err = sql.Open("sqlite3", "database/sqlite-database.db")

	if err != nil {
		log.Fatal(err)
	}

	return err
}
