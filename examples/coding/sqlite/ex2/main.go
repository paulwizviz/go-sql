package main

import (
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	createTblSQL = `CREATE TABLE IF NOT EXISTS t(id INTEGER PRIMARY KEY);`
	insertTblSQL = `INSERT INTO t DEFAULT VALUES;`
	selectAllSQL = `SELECT id FROM t;`
)

func main() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatalf("Unable to get db connection: %v", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	if _, err := db.Exec(createTblSQL); err != nil {
		log.Fatalf("Unable to create table: %v", err)
	}

	committed := false

	txn, err := db.Begin()
	if err != nil {
		log.Fatalf("Unable to start a transaction: %v", err)
	}

	defer func() {
		if !committed {
			txn.Rollback()
		}
	}()

	stmt, err := txn.Prepare(insertTblSQL)
	if err != nil {
		log.Fatalf("Unable to prepare statement: %v", err)
	}
	defer stmt.Close()

	r, err := stmt.Exec()
	if err != nil {
		log.Fatalf("Unable to execute statement: %v", err)
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("First transaction ", id)

	r, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Unable to execute statement: %v", err)
	}

	id, err = r.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Second transaction ", id)

	if err := txn.Commit(); err != nil {
		log.Fatalf("Unable to commit transaction: %v", err)
	}

	committed = true

	rows, err := db.Query(selectAllSQL)
	if err != nil {
		log.Fatalf("Unable to query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		log.Println("Row id:", id)
	}

}
