package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-sql/internal/sqlops"

	_ "github.com/mattn/go-sqlite3"
)

var (
	createTableSQL = "CREATE TABLE IF NOT EXISTS lottery(id INTEGER PRIMARY KEY, ball1 INTEGER, ball2 INTEGER)"
	insertSQL      = "INSERT INTO lottery (ball1, ball2) VALUES (?,?)"
	selectSQL      = "SELECT * FROM lottery WHERE ball1=? AND ball2=?"
	dropTableSQL   = "DROP TABLE lottery"
)

func execInsertStmt(stmt *sql.Stmt, args []int) error {
	r, err := stmt.Exec(args[0], args[1])
	if err != nil {
		return err
	}
	id, err := r.LastInsertId()
	if err != nil {
		log.Printf("ID Error: %v", err)
	}
	rows, err := r.RowsAffected()
	if err != nil {
		log.Printf("Rows Error: %v", err)
	}
	log.Printf("Last insert ID: %v Rows affected: %v", id, rows)
	return nil
}

func execSelectQuery(stmt *sql.Stmt, arg1, arg2 int) error {
	rows, err := stmt.Query(arg1, arg2)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id, ball1, ball2 int
	for rows.Next() {
		err := rows.Scan(&id, &ball1, &ball2)
		if err != nil {
			log.Printf("Error: %v", err)
		}
		fmt.Println(id, ball1, ball2)
	}
	return nil
}

func main() {

	// Instantiate a DB server in memory
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatalf("Connect Error: %v", err)
	}

	// Creating Tables
	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Create table error: %v", err)
	}

	// Create statement to insert into table
	stmt1, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatalf("Prepare insert stmt error: %v", err)
	}

	// Execute Insert Statement to inset values
	// 1 and 2
	if err := execInsertStmt(stmt1, []int{1, 2}); err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	// Execute Insert Statement to inset values
	// 1 and 2 again
	if err := execInsertStmt(stmt1, []int{1, 2}); err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	if err := stmt1.Close(); err != nil {
		log.Fatal(err)
	}

	// Prepare query statment to select result when
	// two values are matched
	stmt2, err := db.Prepare(selectSQL)
	if err != nil {
		log.Fatalf("Prepare select stmt error: %v", err)
	}

	// Execute query statement
	if err := execSelectQuery(stmt2, 1, 2); err != nil {
		log.Fatalf("Select query error: %v", err)
	}

	if err := stmt2.Close(); err != nil {
		log.Fatal(err)
	}

	// Execute statement to drop tables
	if _, err := db.Exec(dropTableSQL); err != nil {
		log.Fatalf("Drop table error: %v", err)
	}
}
