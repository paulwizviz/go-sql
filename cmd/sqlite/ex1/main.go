package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./ex1.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	rs, err := stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

	stmt, err = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

	rs, err = stmt.Exec("Nic", "Raboy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rs)

	rows, err := db.Query("SELECT id, firstname, lastname FROM people")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var fn string
	var ln string
	for rows.Next() {
		rows.Scan(&id, &fn, &ln)
		fmt.Println(strconv.Itoa(id) + ": " + fn + " " + ln)
	}
}
