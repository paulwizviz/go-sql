package main

import (
	"database/sql"
	"fmt"
	"go-sql/internal/pg"
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
	db, err := pg.NewDB("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(createTableStmtStr); err != nil {
		log.Fatalf("Create table error: %v", err)
	}

	stmt1, err := db.Prepare(insertStmtStr)
	if err != nil {
		log.Fatalf("Prepare insert stmt error: %v", err)
	}
	defer stmt1.Close()

	err = insertStatement(stmt1, []int{1, 2})
	if err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	stmt2, err := db.Prepare(selectStmtStr)
	if err != nil {
		log.Fatalf("Prepare select stmt error: %v", err)
	}
	defer stmt2.Close()

	err = selectQuery(stmt2, 1, 2)
	if err != nil {
		log.Fatalf("Select query error: %v", err)
	}

	if _, err := db.Exec(dropTableStmtStr); err != nil {
		log.Fatalf("Drop table error: %v", err)
	}
}
