package sqlops

import (
	"context"
	"database/sql"
)

// TblCreator is an interface to database creation
type TblCreator interface {
	Create(context.Context, *sql.DB) error
}

// TblCreatorFunc is a functional implementation of SQLTableCreator interface
type TblCreatorFunc func(context.Context, *sql.DB) error

func (s TblCreatorFunc) Create(ctx context.Context, db *sql.DB) error {
	return s(ctx, db)
}

// CreateTable is an operation to create table table
func CreateTable(ctx context.Context, creator TblCreator, db *sql.DB) error {
	return creator.Create(ctx, db)
}
