package main

import (
	"database/sql"
	"fmt"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/lib/pq"
)

var (
	createTableStmtStr = "CREATE TABLE IF NOT EXISTS lottery(ball1 INT, ball2 INT)"
	insertStmtStr      = "INSERT INTO lottery (ball1, ball2) VALUES ($1,$2)"
	selectStmtStr      = "SELECT * FROM lottery WHERE ball1=$1 AND ball2=$2"
	dropTableStmtStr   = "DROP TABLE lottery"
)

func insertStatement(stmt *sql.Stmt, args []int) error {
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

func selectQuery(stmt *sql.Stmt, arg1, arg2 int) error {
	rows, err := stmt.Query(arg1, arg2)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ball1, ball2 int
	for rows.Next() {
		err := rows.Scan(&ball1, &ball2)
		if err != nil {
			log.Printf("Error: %v", err)
		}
		fmt.Println(ball1, ball2)
	}
	return nil
}

func main() {

	// STEP 1: Establish connection to Postgres server
	db, err := sqlops.NewPGConn("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}

	// STEP 6: Close connection
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// STEP 2: Create Table
	if _, err := db.Exec(createTableStmtStr); err != nil {
		log.Fatalf("Create table error: %v", err)
	}

	// STEP 3: Insert statement
	// STEP 3a: Prepare insert statement
	stmt1, err := db.Prepare(insertStmtStr)
	if err != nil {
		log.Fatalf("Prepare insert stmt error: %v", err)
	}

	// STEP 3b: Execute insert statement
	if err = insertStatement(stmt1, []int{1, 2}); err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	if err := stmt1.Close(); err != nil {
		log.Fatal(err)
	}

	// STEP 4: SELECT statement
	// STEP 4a: Prepare select statement
	stmt2, err := db.Prepare(selectStmtStr)
	if err != nil {
		log.Fatalf("Prepare select stmt error: %v", err)
	}
	// STEP 4b: Execute select statement
	if err := selectQuery(stmt2, 1, 2); err != nil {
		log.Fatalf("Select query error: %v", err)
	}

	if err := stmt2.Close(); err != nil {
		log.Fatal(err)
	}

	// STEP 5: Drop Table
	if _, err := db.Exec(dropTableStmtStr); err != nil {
		log.Fatalf("Drop table error: %v", err)
	}
}
