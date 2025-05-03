package sqlops

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrConn  = errors.New("db connect error")
	ErrStmt  = errors.New("statement error")
	ErrTable = errors.New("create table error")
)

type Entity any

type SQLTableCreator interface {
	Create(db *sql.DB, entities ...Entity) *sql.Stmt
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

func Ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}
