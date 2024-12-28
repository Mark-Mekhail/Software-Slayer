package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

const MAX_DB_OPEN_RETRIES = 5

func Open() {
	config := mysql.Config{
		User: "software-slayer",
		Passwd: "software-slayer-password",
		Net: "tcp",
		Addr: "mysql:3306",
		DBName: "software-slayer-db",
	}

	var err error
	for i := 0; i < MAX_DB_OPEN_RETRIES; i++ {
		db, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	db.Close()
}

func Exec(query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	return result, err
}

func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	return rows, err
}

func QueryRow(query string, args ...any) *sql.Row {
	return db.QueryRow(query, args...)
}