// package main
package config

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Declare a global database variable
var DB *sql.DB

// Function to initialize the database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "users.db")
	if err != nil {
		panic(err)
	}

	// Create new table if doesn't ecist
	query := `CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT NOT NULL, password TEXT NOT NULL);`
	_, err = DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database initialized!")
}

// func main() {
// 	initDB()
// 	fmt.Println("Database initialized!")  // working fine
// }
