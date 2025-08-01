package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./issues.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	createTable()
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS issues (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					title TEXT NOT NULL,
					content TEXT NOT NULL,
					status INTEGER NOT NULL,
					severity TEXT NOT NULL
					);`
	if _, err := DB.Exec(query); err != nil {
		log.Fatal(err)
	}
}
