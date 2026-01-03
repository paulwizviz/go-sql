package main

import (
	"fmt"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	// STEP 1: Connect to Postgres
	db, err := sqlops.NewPGConn("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}

	// STEP 3: Close DB connection
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// STEP 2: Ping Postgres server
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB pinged")
}
