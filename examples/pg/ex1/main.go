package main

import (
	"fmt"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlops.NewPGConn("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinging DB")
}
