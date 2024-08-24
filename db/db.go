package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(250) NOT NULL,
		description TEXT NOT NULL,
		location VARCHAR(250) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create events table:", err)
	}
}
