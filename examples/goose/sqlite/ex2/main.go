package main

import (
	"context"
	"log"

	"github.com/pressly/goose/v3"

	_ "modernc.org/sqlite"
)

func main() {

	db, err := goose.OpenDBWithDriver("sqlite", "./sqlite.db")
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v", err)
		}
	}()

	if err := goose.RunContext(context.Background(), "up", db, "."); err != nil {
		log.Fatalf("goose: migration failed: %v", err)
	}
}
