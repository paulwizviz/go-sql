package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrConn  = errors.New("db connect error")
	ErrStmt  = errors.New("statement error")
	ErrTable = errors.New("create table error")
)

const (
	ver = "sqlite3"
)

func NewMemDB() (*sql.DB, error) {
	db, err := sql.Open(ver, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func NewDBFile(f string) (*sql.DB, error) {
	db, err := sql.Open(ver, f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}
