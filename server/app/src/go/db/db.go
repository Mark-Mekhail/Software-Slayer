package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"

	"software-slayer/configs"
)

// Database wraps and extends the standard sql.DB functionality
type Database struct {
	conn *sql.DB
}

// OpenConnection establishes a connection to the MySQL database with retries and configures the connection pool
func OpenConnection(user, password, address, dbName string) (*sql.DB, error) {
	config := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 address,
		DBName:               dbName,
		AllowNativePasswords: true,
		ParseTime:            true,
		Timeout:              5 * time.Second,
	}

	dsn := config.FormatDSN()
	log.Printf("Connecting to database at %s...", address)

	// Retry connection with exponential backoff
	var conn *sql.DB
	var err error
	retryDelay := time.Second

	for i := 0; i < configs.MAX_DB_OPEN_RETRIES; i++ {
		if i > 0 {
			log.Printf("Retrying connection to database (attempt %d/%d)", i+1, configs.MAX_DB_OPEN_RETRIES)
			time.Sleep(retryDelay)
			// Double the delay for next retry, up to a maximum of 16 seconds
			if retryDelay < 16*time.Second {
				retryDelay *= 2
			}
		}

		conn, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			continue
		}

		// Test the connection
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err = conn.PingContext(ctx); err != nil {
			log.Printf("Failed to ping database: %v", err)
			conn.Close()
			continue
		}

		// Successfully connected
		break
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w",
			configs.MAX_DB_OPEN_RETRIES, err)
	}

	// Configure connection pool
	conn.SetMaxOpenConns(25)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(5 * time.Minute)
	conn.SetConnMaxIdleTime(5 * time.Minute)

	log.Printf("Successfully connected to database at %s", address)
	return conn, nil
}

// NewDB creates a new Database instance
func NewDB(conn *sql.DB) *Database {
	return &Database{conn: conn}
}

// Close closes the database connection
func (db *Database) Close() error {
	log.Println("Closing database connection")
	return db.conn.Close()
}

// ExecContext executes a query without returning any rows, with a context
func (db *Database) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.conn.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows, with a context
func (db *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.conn.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that returns at most one row, with a context
func (db *Database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return db.conn.QueryRowContext(ctx, query, args...)
}
