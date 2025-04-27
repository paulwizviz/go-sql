package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-sql/internal/sqlite"
)

var (
	createTableStmtStr = "CREATE TABLE IF NOT EXISTS lottery(ball1 INT, ball2 INT)"
	insertStmtStr      = "INSERT INTO lottery (ball1, ball2) VALUES (?,?)"
	selectStmtStr      = `SELECT (SELECT ball1 FROM lottery WHERE ball1 =? OR ball1 =?) AS ball1, 
	                        (SELECT ball2 FROM lottery WHERE ball2=? OR ball2=?) AS ball2 
					FROM lottery`
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
	rows, err := stmt.Query(arg1, arg2, arg1, arg2)
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
	db, err := sqlite.ConnectMemDefault()
	if err != nil {
		log.Fatalf("Connect Error: %v", err)
	}

	err = sqlite.CreateTable(db, createTableStmtStr)
	if err != nil {
		log.Fatalf("Create Table error: %v", err)
	}

	stmt1, err := sqlite.PrepareStatement(db, insertStmtStr)
	if err != nil {
		log.Fatalf("Prepare insert stmt error: %v", err)
	}
	defer stmt1.Close()

	err = insertStatement(stmt1, []int{1, 2})
	if err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	stmt2, err := sqlite.PrepareStatement(db, selectStmtStr)
	if err != nil {
		log.Fatalf("Prepare select stmt error: %v", err)
	}
	defer stmt2.Close()

	err = selectQuery(stmt2, 1, 2)
	if err != nil {
		log.Fatalf("Select query error: %v", err)
	}
}
