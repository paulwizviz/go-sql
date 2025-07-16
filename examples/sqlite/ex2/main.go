package main

import (
	"context"
	"go-sql/internal/person"
	"go-sql/internal/sqlops"
	"log"
)

func main() {

	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatal(err)
	}

	person.CreateTables(context.TODO(), db)

}
