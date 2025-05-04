package main

import (
	"fmt"
	"go-sql/internal/pg"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := pg.NewDB("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinging DB")
}
