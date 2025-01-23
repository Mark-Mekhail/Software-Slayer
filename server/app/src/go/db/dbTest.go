package db

import (
	"database/sql"
)

func NewTestDB(testDB *sql.DB) *Database {
	return &Database{conn: testDB}
}