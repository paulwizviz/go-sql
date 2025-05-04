package main

import (
	"fmt"
	"go-sql/internal/pg"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := pg.NewDB("postgres", "postgres", "localhost", 5432, "default")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = sqlops.Ping(db)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Pinging DB")
}
