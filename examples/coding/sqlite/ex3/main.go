package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const showTables = "SELECT name FROM sqlite_master WHERE type='table';"

func createTable(ctx context.Context, db *sql.DB) error {
	if err := sqlops.CreateTableTx(ctx,
		db,
		CreateTblPersonSQLFn,
		CreateTblPNIFn,
		CreateTblNameIdentifierSQLFn,
	); err != nil {
		return err
	}

	rows, err := db.Query(showTables)
	if err != nil {
		return err
	}

	for rows.Next() {
		var tblName string
		rows.Scan(&tblName)
		log.Println(tblName)
	}

	return nil
}

func writePerson(ctx context.Context, db *sql.DB, data []any) ([]any, error) {
	return sqlops.Writer(ctx, db, createTblNameIdentifierSQL, data, InsertPersonFnc)
}

func main() {
	// Instantiate SQLite DB in memory
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer db.Close()

	ctx := context.TODO()
	p := Person{}
	p1 := Person{}
	createTable(ctx, db)
	r, err := writePerson(ctx, db, []any{p, p1})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
