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

func OpenConnection(user, password, address, dbName string) (*sql.DB, error) {
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

	return conn, nil
}

func NewDB(conn *sql.DB) *Database {
	return &Database{conn: conn}
}

func (db *Database) Close() error {
	return db.conn.Close()
}

func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	return db.conn.Exec(query, args...)
}

func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	return db.conn.Query(query, args...)
}

func (db *Database) QueryRow(query string, args ...any) *sql.Row {
	return db.conn.QueryRow(query, args...)
}
