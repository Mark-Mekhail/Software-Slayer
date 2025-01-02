package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

const MAX_DB_OPEN_RETRIES = 5

var db *sql.DB

/*
 * Open opens a connection to the database and initializes the db variable.
 */
func Open(user, password, address, dbName string) {
	config := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   address,
		DBName: dbName,
	}

	// Retry opening the database connection a few times before giving up
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

	// Ping the database to ensure the connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

/*
 * Close closes the connection to the database.
 */
func Close() {
	db.Close()
}

/*
 * Exec executes a query that does not return rows.
 * @param query: The query to execute.
 * @param args: The arguments to pass to the query.
 * @return sql.Result: The result of the query.
 * @return error: An error if the query could not be executed.
 */
func Exec(query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	return result, err
}

/*
 * Query executes a query that returns rows.
 * @param query: The query to execute
 * @param args: The arguments to pass to the query
 * @return *sql.Rows: The rows returned by the query
 * @return error: An error if the query could not be executed
 */
func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	return rows, err
}

/*
 * QueryRow executes a query that returns a single row.
 * @param query: The query to execute
 * @param args: The arguments to pass to the query
 * @return *sql.Row: The row returned by the query
 */
func QueryRow(query string, args ...any) *sql.Row {
	return db.QueryRow(query, args...)
}
