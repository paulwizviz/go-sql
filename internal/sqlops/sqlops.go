package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrDB    = errors.New("DB error")
	ErrStmt  = errors.New("statement error")
	ErrTable = errors.New("create table")
	ErrTxn   = errors.New("transaction error")
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
		return fmt.Errorf("%w: begin transaction failed: %w", ErrTxn, err)
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
			return fmt.Errorf("%w: creating table: %w", ErrTable, err)
		}
	}

	// The return statement is simplified, returning the result of tx.Commit().
	// This correctly handles and propagates the commit error.
	return tx.Commit()
}

type StmtWriter struct {
	stmt *sql.Stmt
}

func (s *StmtWriter) Exec(ctx context.Context, args ...any) error {
	_, err := s.stmt.ExecContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("%w: executing statement: %w", ErrStmt, err)
	}
	return nil
}

func (s *StmtWriter) Close() error {
	err := s.stmt.Close()
	if err != nil {
		return fmt.Errorf("%w: insert statement closing: %w", ErrStmt, err)
	}
	return nil
}

type WriteStmtArg struct {
	stmt *sql.Stmt
	args []any
}

func ExecuteWriteStmtTx(ctx context.Context, db *sql.DB, wsas ...WriteStmtArg) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: unable to instantiat transaction: %w", ErrTxn, err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, wsa := range wsas {
		_, err := tx.StmtContext(ctx, wsa.stmt).ExecContext(ctx, wsa.args...)
		if err != nil {
			return fmt.Errorf("%w: unable to execute: %w", ErrTxn, err)
		}
	}

	return tx.Commit()
}

// RowProcessor is a callback function for processing a single database row.
type RowProcessor func(rows *sql.Rows) error

// StmtReader encapsulates a prepared statement for SELECT queries.
type StmtReader struct {
	stmt *sql.Stmt
}

// Query executes the prepared statement and processes each row using a callback.
// The library handles the rows.Close() call, simplifying usage for the user.
func (s *StmtReader) Query(ctx context.Context, processor RowProcessor, args ...any) error {
	rows, err := s.stmt.QueryContext(ctx, args...)
	if err != nil {
		return fmt.Errorf("%w: executing query: %w", ErrStmt, err)
	}
	defer rows.Close() // The library is now responsible for closing rows.

	for rows.Next() {
		if err := processor(rows); err != nil {
			return fmt.Errorf("%w: row processing failed: %w", ErrStmt, err)
		}
	}

	return rows.Err() // Check for any error that may have occurred during iteration.
}

// ReadOp holds a prepared statement and a callback to process its results.
type ReadOp struct {
	Stmt      *sql.Stmt
	Args      []any
	Processor RowProcessor
}

// ExecuteReadOpTx executes multiple SELECT statements in a transaction
// and processes each result set with a provided callback.
func ExecuteReadOpTx(ctx context.Context, db *sql.DB, ops ...ReadOp) (err error) {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w: begin transaction failed: %w", ErrTxn, err)
	}

	// Named return 'err' is used to correctly handle deferred rollback.
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, op := range ops {
		rows, err := tx.StmtContext(ctx, op.Stmt).QueryContext(ctx, op.Args...)
		if err != nil {
			return fmt.Errorf("%w: query execution failed: %w", ErrStmt, err)
		}

		// This is the key change: The library handles rows.Next() and rows.Close().
		for rows.Next() {
			if err = op.Processor(rows); err != nil {
				rows.Close() // Close rows on processing error
				return fmt.Errorf("%w: row processing failed: %w", ErrStmt, err)
			}
		}

		if err = rows.Close(); err != nil {
			return fmt.Errorf("%w: failed to close rows: %w", ErrStmt, err)
		}
	}

	return tx.Commit()
}
