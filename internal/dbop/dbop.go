package dbop

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrPrepareStatement = errors.New(("unable to prepare statement"))
	ErrTableCreation    = errors.New("unable to create table")
)

type ValidCreateTblStmtStrFn func(string) error

var (
	ValidCreateTblStmtStr ValidCreateTblStmtStrFn = func(sstmt string) error {
		// TODO add parser here
		return nil
	}
)

func CreateTable(db *sql.DB, sstmt string, fn ValidCreateTblStmtStrFn) error {
	err := fn(sstmt)
	if err != nil {
		return err
	}
	_, err = db.Exec(sstmt)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrTableCreation, err.Error())
	}
	return nil
}
