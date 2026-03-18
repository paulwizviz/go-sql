package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
)

var (
	ErrDB          = errors.New("db error")
	ErrCreateTable = errors.New("create table error")
	ErrPrepareStmt = errors.New("prepare statement error")
	ErrTx          = errors.New("transaction error")
)

// TblCreatorTxFunc is the idiomatic way to handle a function callback.
// The interface definition can be omitted if only functions will be used.
type TblCreatorTxFunc func(context.Context, *sql.Tx) error

// CreateTableTx is an operation to create tables over a transaction.
// Using a named return variable 'err' is crucial for the deferred rollback.
func CreateTableTx(ctx context.Context, db *sql.DB, creators ...TblCreatorTxFunc) (err error) {

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrTx, err)
	}

	// This defer correctly captures the named return variable 'err'.
	defer func() {
		if err != nil {
			// A rollback is performed only if the function is exiting with an error.
			tx.Rollback()
		}
	}()

	for _, creator := range creators {
		// Assigning the error to the named return variable 'err'.
		if err = creator(ctx, tx); err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTable, err)
		}
	}

	// The return statement is simplified, returning the result of tx.Commit().
	// This correctly handles and propagates the commit error.
	return tx.Commit()
}

type RowWriterFunc func(context.Context, *sql.Tx, *sql.Stmt, any) (any, error)

func Writer(ctx context.Context, db *sql.DB, rawStmt string, dataList []any, rowWriter RowWriterFunc) ([]any, error) {
	if len(dataList) == 0 {
		return nil, nil
	}

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTx, err)
	}

	// This defer correctly captures the named return variable 'err'.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, rawStmt)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrPrepareStmt, err)
	}
	defer stmt.Close()

	dataSet := []any{}
	for _, data := range dataList {
		d, err := rowWriter(ctx, tx, stmt, data)
		if err != nil {
			slog.Info(err.Error())
			continue
		}
		dataSet = append(dataSet, d)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrTx, err)
	}

	return dataSet, nil
}
