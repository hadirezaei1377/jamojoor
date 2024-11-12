package pdfexelservice

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

// InitDB initializes the SQLite database
func InitDB() {
	var err error
	db, err = sqlx.Connect("sqlite3", "users.db")
	if err != nil {
		log.Fatalln(err)
	}

	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name  TEXT,
        email TEXT
    );
    `
	db.MustExec(schema)
}
