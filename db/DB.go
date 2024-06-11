package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	CreateTables()
}

func CreateTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime DATETIME NOT NULL,
        user_id INTEGER
    );`
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create DB table: ", err)
	}
}
