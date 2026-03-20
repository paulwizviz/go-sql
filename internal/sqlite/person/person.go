package person

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go-sql/internal/sqlops"
)

var (
	ErrCreateTbl = errors.New("unable to create person table")
)

const (
	// Person Table
	tblPerson = "person"
	colPerID  = "id"

	// Name Identifier Table
	tblNameIdentifier = "name_identifier"
	colNameID         = "id"
	colNameIDFirst    = "first_name"
	colNameIDMiddle   = "middle_name"
	colNameIDSurname  = "surname"

	// Person Name Identifier Table
	tblPNI         = "person_name_identifier"
	colPNIID       = "id"
	colPNINameID   = "name_identifier_id"
	colPNIPersonID = "person_id"
)

var (
	// Person Table
	createTblPersonSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
%[2]s INTEGER PRIMARY KEY 
	);`, tblPerson, colPerID)

	CreateTblPersonSQLFn sqlops.TblCreatorTxFunc = func(ctx context.Context, txn *sql.Tx) error {
		_, err := txn.ExecContext(ctx, createTblPersonSQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTbl, err)
		}
		return nil
	}

	insertPersonSQL = fmt.Sprintf(`INSERT INTO %[1]s DEFAULT VALUES;`, tblPerson)

	// Name Identifier Table
	createTblNameIDSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
	%[2]s INTEGER PRIMARY KEY,
	%[3]s TEXT NOT NULL,
	%[4]s TEXT,
	%[5]s TEXT NOT NULL
);`, tblNameIdentifier, colNameID, colNameIDFirst, colNameIDMiddle, colNameIDSurname)

	CreateTblNameIDSQLFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, createTblNameIDSQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTbl, err)
		}
		return nil
	}

	// Person Name Identifier Table
	createTblPNISQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %[1]s (
	%[2]s INTEGER PRIMARY KEY,
	%[3]s INTEGER NOT NULL,
	%[4]s INTEGER NOT NULL,
	FOREIGN KEY (%[3]s) REFERENCES %[5]s(%[6]s),
	FOREIGN KEY (%[4]s) REFERENCES %[7]s(%[8]s)  
	);`, tblPNI, colPNIID, colPNINameID, colPNIPersonID, tblNameIdentifier, colNameID, tblPerson, colPerID)

	CreateTblPNIFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, createTblPNISQL)
		if err != nil {
			return fmt.Errorf("%w: %v", ErrCreateTbl, err)
		}
		return nil
	}
)

func insertPerTbl(ctx context.Context, txn *sql.Tx) (int64, error) {
	result, err := txn.ExecContext(ctx, insertPersonSQL)
	if err != nil {
		return 0, fmt.Errorf("Error inserting person: %v", err)
	}
	return result.LastInsertId()
}

func insertNITbl(ctx context.Context, txn *sql.Tx, firstName, middleName, surname string) (int64, error) {
	result, err := txn.ExecContext(ctx, fmt.Sprintf(`INSERT INTO %[1]s (%[2]s, %[3]s, %[4]s) VALUES (?, ?, ?);`, tblNameIdentifier, colNameIDFirst, colNameIDMiddle, colNameIDSurname), firstName, middleName, surname)
	if err != nil {
		return 0, fmt.Errorf("Error inserting name identifier: %v", err)
	}
	return result.LastInsertId()
}

func insertPNITbl(ctx context.Context, txn *sql.Tx, personID, nameID int64) (int64, error) {
	result, err := txn.ExecContext(ctx, fmt.Sprintf(`INSERT INTO %[1]s (%[2]s, %[3]s) VALUES (?, ?);`, tblPNI, colPNINameID, colPNIPersonID), nameID, personID)
	if err != nil {
		return 0, fmt.Errorf("Error inserting person-name identifier: %v", err)
	}
	return result.LastInsertId()
}

type Detail struct {
	ID         int64
	FirstName  string
	MiddleName sql.NullString
	Surname    string
}

func PersistPersonData(ctx context.Context, db *sql.DB, firstName, middleName, surname string) (Detail, error) {

	txn, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	// This defer correctly captures the named return variable 'err'.
	defer func() {
		if err != nil {
			// A rollback is performed only if the function is exiting with an error.
			txn.Rollback()
		}
	}()

	personID, err := insertPerTbl(ctx, txn)
	if err != nil {
		return Detail{}, err
	}

	nameID, err := insertNITbl(ctx, txn, firstName, middleName, surname)
	if err != nil {
		return Detail{}, err
	}

	pniID, err := insertPNITbl(ctx, txn, personID, nameID)
	if err != nil {
		return Detail{}, err
	}

	txn.Commit()

	return Detail{
		ID:        pniID,
		FirstName: firstName,
		MiddleName: sql.NullString{
			String: middleName,
			Valid:  middleName != "",
		},
		Surname: surname,
	}, nil
}
