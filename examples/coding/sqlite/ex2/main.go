package main

import (
	"context"
	"fmt"
	"go-sql/internal/person"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Instantiate SQLite DB
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer db.Close()

	ctx := context.Background()

	// SQLite create table
	err = person.SQLiteCreateTable(ctx, db)
	if err != nil {
		log.Fatalf("Create table error: %v", err)
	}

	err = person.SQLiteInsertName(ctx, db)
	if err != nil {
		log.Fatalf("Insert name error: %v", err)
	}

	pi, err := person.GetNames(ctx, db)
	if err != nil {
		log.Fatalf("get names: %v", err)
	}

	fmt.Println(pi)
}
