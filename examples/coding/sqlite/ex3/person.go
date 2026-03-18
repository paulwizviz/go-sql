package main

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"go-sql/internal/sqlops"
)

var (
	ErrTblPerson       = errors.New("unable to create person table")
	ErrTblNamedID      = errors.New("unable to create named_identifier table")
	ErrTblPersonNameID = errors.New("unable to create person_name_identifier table")
)

const (
	tblPerson = "person"
	colPerID  = "id"
)

var (
	ErrCreatePersonTbl = errors.New("create person table")
	ErrInsertPerson    = errors.New("insert into person")

	createTblPersonSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
%[2]s INTEGER PRIMARY KEY 
	);`, tblPerson, colPerID)

	CreateTblPersonSQLFn sqlops.TblCreatorTxFunc = func(ctx context.Context, txn *sql.Tx) error {
		_, err := txn.ExecContext(ctx, createTblPersonSQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreatePersonTbl, err)
		}
		return nil
	}

	InsertPersonSQL = fmt.Sprintf(`INSERT INTO %[1]s DEFAULT VALUES;`, tblPerson)

	InsertPersonFnc sqlops.RowWriterFunc = func(ctx context.Context, txn *sql.Tx, stmt *sql.Stmt, data any) (any, error) {
		r, err := stmt.Exec()
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInsertPerson, err)
		}
		d := data.(Person)
		d.ID, err = r.LastInsertId()
		if err != nil {
			return nil, err
		}
		return d, nil
	}
)

const (
	tblNameIdentifier = "name_identifier"
	colNameID         = "id"
	colNameIDFirst    = "first_name"
	colNameIDMiddle   = "middle_name"
	colNameIDSurname  = "surname"
)

var (
	ErrCreateNameIdentifierTbl = errors.New("create name identifier table")

	createTblNameIdentifierSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
	%[2]s INTEGER PRIMARY KEY,
	%[3]s TEXT NOT NULL,
	%[4]s TEXT,
	%[5]s TEXT NOT NULL
);`, tblNameIdentifier, colNameID, colNameIDFirst, colNameIDMiddle, colNameIDSurname)

	CreateTblNameIdentifierSQLFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, createTblNameIdentifierSQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateNameIdentifierTbl, err)
		}
		return nil
	}
)

const (
	tblPersonNameIdentifier = "person_name_identifier"
	colPNIID                = "id"
	colPNINameID            = "name_identifier_id"
	colPNIPersonID          = "person_id"
)

var (
	ErrCreatePNITbl = errors.New("create person name identifier table")

	createTblPNISQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
	%[2]s INTEGER PRIMARY KEY,
	%[3]s INTEGER NOT NULL,
	%[4]s INTEGER NOT NULL,
	FOREIGN KEY (%[3]s) REFERENCES %[5]s(%[6]s),
	FOREIGN KEY (%[4]s) REFERENCES %[7]s(%[8]s)  
	);`, tblPersonNameIdentifier, colPNIID, colPNINameID, colPNIPersonID, tblNameIdentifier, colNameID, tblPerson, colPerID)

	CreateTblPNIFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, createTblPNISQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateNameIdentifierTbl, err)
		}
		return nil
	}
)

// NameValue is a type of object value
type NameValue string

type Person struct {
	ID        int64     `json:"person_id"`
	FirstName NameValue `json:"first_name"`
	Surname   NameValue `json:"middlename"`
	Nickname  NameValue `json:"nickname"`
}
