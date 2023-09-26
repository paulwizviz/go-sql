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
	ErrTable = errors.New("creating table error")
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

func ConnectMemPool(f string, maxPool int) (*sql.DB, error) {
	db, err := sql.Open(ver, fmt.Sprintf("Data Source=%s;Pooling=True;Max Pool Size=%d", f, maxPool))
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}
