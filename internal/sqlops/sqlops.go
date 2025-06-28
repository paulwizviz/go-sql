package sqlops

import (
	"database/sql"
)

// SQLTableCreator is an interface to database creation
type SQLTableCreator interface {
	Create(*sql.DB) error
}

// SQLTableCreatorFunc is a functional implementation of SQLTableCreator interface
type SQLTableCreatorFunc func(*sql.DB) error

func (s SQLTableCreatorFunc) Create(db *sql.DB) error {
	return s(db)
}

// CreateTable is an operation to create table table
func CreateTable(creator SQLTableCreator, db *sql.DB) error {
	return creator.Create(db)
}
