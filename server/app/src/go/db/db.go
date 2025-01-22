package db

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"

	"software-slayer/configs"
)

type Database struct {
	conn *sql.DB
}

type DatabaseInterface interface {
	Close() error
	Exec(query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
}

/*
 * newDB creates a new connection to the database.
 * @param user: The username to use to connect to the database.
 * @param password: The password to use to connect to the database.
 * @param address: The address of the database.
 * @param dbName: The name of the database.
 * @return *DB: A new connection to the database.
 * @return error: An error if the connection could not be established.
 */
func NewDB(user, password, address, dbName string) (*Database, error) {
	config := mysql.Config{
		User:   user,
		Passwd: password,
		Net:    "tcp",
		Addr:   address,
		DBName: dbName,
	}

	// Retry opening the database connection a few times before giving up
	var conn *sql.DB
	var err error
	for i := 0; i < configs.MAX_DB_OPEN_RETRIES; i++ {
		conn, err = sql.Open("mysql", config.FormatDSN())
		if err != nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	// Ping the database to ensure the connection is valid
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{conn: conn}, nil
}

/*
 * Close closes the connection to the database.
 */
func (db *Database) Close() error {
	return db.conn.Close()
}

/*
 * Exec executes a query that does not return rows.
 * @param query: The query to execute.
 * @param args: The arguments to pass to the query.
 * @return sql.Result: The result of the query.
 * @return error: An error if the query could not be executed.
 */
func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

/*
 * Query executes a query that returns rows.
 * @param query: The query to execute
 * @param args: The arguments to pass to the query
 * @return *sql.Rows: The rows returned by the query
 * @return error: An error if the query could not be executed
 */
func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

/*
 * QueryRow executes a query that returns a single row.
 * @param query: The query to execute
 * @param args: The arguments to pass to the query
 * @return *sql.Row: The row returned by the query
 */
func (db *Database) QueryRow(query string, args ...any) *sql.Row {
	return db.conn.QueryRow(query, args...)
}
