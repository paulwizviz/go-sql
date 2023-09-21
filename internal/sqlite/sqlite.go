package sqlite

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Connect(maxIdleTm time.Duration, maxLifeTm time.Duration, maxIdleConn int, maxOpenConn int) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(maxIdleTm)
	db.SetConnMaxLifetime(maxLifeTm)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)
	return db, nil
}

func CreateTable(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
