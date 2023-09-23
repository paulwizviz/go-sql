package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrConn = errors.New("db connect error")
	ErrStmt = errors.New("statement error")
)

type Config struct {
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	MaxIdleConn     int
	MaxOpenConn     int
}

func ConnectMem(config Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnMaxLifeTime)
	db.SetMaxIdleConns(config.MaxIdleConn)
	db.SetMaxOpenConns(config.MaxOpenConn)
	return db, nil
}

func CreateTable(db *sql.DB, strStmt string) error {
	stmt, err := db.Prepare(strStmt)
	if err != nil {
		return fmt.Errorf("%w-%v", ErrStmt, err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
