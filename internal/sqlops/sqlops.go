package sqlops

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrTxn   = errors.New("transaction error")
	ErrTable = errors.New("create table")
	ErrWrite = errors.New("write table")
)

// TblCreatorTx is an interface to create Table
type TblCreatorTx interface {
	Create(context.Context, *sql.Tx) error
}

// TblCreatorTxFunc is a functional implementation of TblCreatorTx
// interface
type TblCreatorTxFunc func(context.Context, *sql.Tx) error

func (t TblCreatorTxFunc) Create(ctx context.Context, tx *sql.Tx) error {
	return t(ctx, tx)
}

// CreateTableTx is an operation to create tables over a transaction
func CreateTableTx(ctx context.Context, db *sql.DB, creators ...TblCreatorTx) error {

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("%w-%v", ErrTxn, err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	for _, creator := range creators {
		err := creator.Create(ctx, tx)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTable, err)
		}
	}
	tx.Commit()
	return nil
}
