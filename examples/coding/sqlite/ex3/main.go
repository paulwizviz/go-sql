package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql/internal/sqlite/person"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const showTables = "SELECT name FROM sqlite_master WHERE type='table';"

func createTables(ctx context.Context, db *sql.DB) error {
	if err := sqlops.CreateTableTx(
		ctx,
		db,
		person.CreateTblPersonSQLFn,
		person.CreateTblNameIDSQLFn,
		person.CreateTblPNIFn,
	); err != nil {
		return fmt.Errorf("Table creation error: %v", err)
	}

	rows, err := db.QueryContext(ctx, showTables)
	if err != nil {
		return fmt.Errorf("Error querying tables: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return fmt.Errorf("Error scanning table name: %v", err)
		}
		log.Printf("Table: %s", tableName)
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	// Instantiate SQLite DB in memory
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer db.Close()

	ctx := context.TODO()
	if err := createTables(ctx, db); err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	// persist a person
	personDetail, err := person.PersistPersonData(ctx, db, "John", "Doe", "Smith")
	if err != nil {
		log.Fatalf("Error persisting person data: %v", err)
	}
	log.Printf("Persisted person: %+v", personDetail)

	personDetail, err = person.PersistPersonData(ctx, db, "Jane", "Doe", "Smith")
	if err != nil {
		log.Fatalf("Error persisting person data: %v", err)
	}
	log.Printf("Persisted person: %+v", personDetail)

}
