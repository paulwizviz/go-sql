package main

import (
	"context"
	"log"

	"github.com/pressly/goose/v3"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connStmt := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	db, err := goose.OpenDBWithDriver("pgx", connStmt)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v", err)
		}
	}()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("goose: failed to set dialect: %v", err)
	}

	if err := goose.RunContext(context.Background(), "up", db, "."); err != nil {
		log.Fatalf("goose: migration failed: %v", err)
	}
}
