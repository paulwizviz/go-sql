package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./ex2.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec()
	stmt, err = db.Prepare("INSERT INTO people (id, firstname, lastname) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	rs, err := stmt.Exec(1, "A", "B")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

	rs, err = stmt.Exec(1, "A", "C")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

}
