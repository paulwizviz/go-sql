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

func ConnectMemDefault() (*sql.DB, error) {
	db, err := sql.Open(ver, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func ConnectFile(f string) (*sql.DB, error) {
	db, err := sql.Open(ver, f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

func CreateTable(db *sql.DB, stmt string) error {
	_, err := db.Exec(stmt)
	if err != nil {
		return fmt.Errorf("%w-%v", ErrTable, err)
	}
	return nil
}

func PrepareStatement(db *sql.DB, query string) (*sql.Stmt, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("%v-%v", ErrStmt, err)
	}
	return stmt, nil
}
