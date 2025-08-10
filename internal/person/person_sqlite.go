package person

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"go-sql/internal/sqlops"
)

// SQLite
var (
	//go:embed sql/sqlite/tbl_person.sql
	SQLiteTblPersonSQL string
	//go:embed sql/sqlite/tbl_named_id.sql
	SQLiteTblNamedIDSQL string
	//go:embed sql/sqlite/tbl_person_name_id.sql
	SQLiteTblPersonNameIDSQL string
	//go:embed sql/sqlite/insert_person.sql
	SQLiteInsertPersonSQL string
	//go:embed sql/sqlite/insert_named_id.sql
	SQLiteInsertNamedIDSQL string
	//go:embed sql/sqlite/insert_person_name_id.sql
	SQLiteInsertPersonNameIDSQL string

	SQLiteCreatePersonTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblPersonSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPerson, err)
		}
		return nil
	}

	SQLiteCreateNamedIDTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblNamedIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblNamedID, err)
		}
		return nil
	}

	SQLiteCreatePersonNameIDTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblPersonNameIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPersonNameID, err)
		}
		return nil
	}
)

func SQLiteCreateTable(ctx context.Context, db *sql.DB) error {
	return sqlops.CreateTableTx(
		ctx,
		db,
		SQLiteCreatePersonTxFn,
		SQLiteCreateNamedIDTxFn,
		SQLiteCreatePersonNameIDTxFn)
}

func SQLiteInsertName(ctx context.Context, db *sql.DB) error {

	var stmts []*sql.Stmt

	insertPersonStmt, err := db.PrepareContext(ctx, SQLiteInsertPersonSQL)
	if err != nil {
		return err
	}
	defer insertPersonStmt.Close()
	stmts = append(stmts, insertPersonStmt)

	insertNameIDStmt, err := db.PrepareContext(ctx, SQLiteInsertNamedIDSQL)
	if err != nil {
		return err
	}
	defer insertNameIDStmt.Close()
	stmts = append(stmts, insertNameIDStmt)

	insertPersonNameIDStmt, err := db.PrepareContext(ctx, SQLiteInsertPersonNameIDSQL)
	if err != nil {
		return err
	}
	defer insertPersonNameIDStmt.Close()
	stmts = append(stmts, insertPersonNameIDStmt)

	firstName := "John"
	middleName := ""
	surname := "Doe"
	err = insertNames(ctx, db, stmts, firstName, middleName, surname)
	if err != nil {
		return fmt.Errorf("unable to insert name: %v", err)
	}
	return nil
}

func insertNames(
	ctx context.Context,
	db *sql.DB,
	stmt []*sql.Stmt,
	firstName string,
	middleName string,
	surname string) error {

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return fmt.Errorf("execute statements: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	result, err := tx.StmtContext(ctx, stmt[0]).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("execute statement: %v", err)
	}

	personID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("obtain person id: %v", err)
	}

	result, err = tx.StmtContext(ctx, stmt[1]).ExecContext(ctx, firstName, middleName, surname)
	if err != nil {
		return fmt.Errorf("insert name error: %v", err)
	}

	nameID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("obtain named id: %v", err)
	}

	_, err = tx.StmtContext(ctx, stmt[2]).ExecContext(ctx, personID, nameID)
	if err != nil {
		return fmt.Errorf("insert ")
	}

	return tx.Commit()
}
