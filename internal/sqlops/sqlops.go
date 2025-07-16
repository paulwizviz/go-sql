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

func (t TblCreatorFunc) Create(ctx context.Context, db *sql.DB) error {
	return t(ctx, db)
}

// CreateTable is an operation to create table table
func CreateTable(ctx context.Context, creator TblCreator, db *sql.DB) error {
	return creator.Create(ctx, db)
}

// TblWriter is an interface to insert data
type TblWriter interface {
	Write(context.Context, *sql.DB, ...any) error
}

type TblWriterFunc func(context.Context, *sql.DB, ...any) error

func (t TblWriterFunc) Write(ctx context.Context, db *sql.DB, args ...any) error {
	return t(ctx, db, args)
}

func WriteTable(ctx context.Context, db *sql.DB, writer TblWriter, args ...any) error {
	return writer.Write(ctx, db, args...)
}
