package retry

import (
	"database/sql"
	"log"
	"retry-project/models"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB // sqlite

func InitDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "./retry.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		endpoint TEXT,
		payload TEXT,
		retry_count INTEGER,
		last_tried DATETIME,
		status TEXT
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func AddRequest(request models.Request) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO requests(endpoint, payload, retry_count, last_tried, status) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(request.Endpoint, request.Payload, request.RetryCount, request.LastTried, request.Status)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
