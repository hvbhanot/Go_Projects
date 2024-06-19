package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DB is a global variable of type *sql.DB used for connecting and interacting with a SQLite database. It is initialized and configured in the InitDB() function. It is used in various functions for executing SQL queries, retrieving and saving data to the database.
var DB *sql.DB

// InitDB initializes the database connection and sets the maximum number of open and idle connections.
// It creates the necessary tables by calling the createTables function.
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

// createTables creates necessary database tables if they do not exist.
func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create events table.")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
	                                  
	   FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

	registrations := `CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    eventId INTEGER,
    userId INTEGER,
    FOREIGN KEY(userId) REFERENCES users(id),
    FOREIGN KEY(eventId) REFERENCES events(id)
)`
	_, err = DB.Exec(registrations)
	if err != nil {
		panic("Could not create registration table.")
	}
}
