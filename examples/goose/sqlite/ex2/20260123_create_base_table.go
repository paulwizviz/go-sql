package main

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	// Register the migration functions with a unique name/version
	goose.AddMigrationContext(UpCreateUsers, DownCreateUsers)
}

func UpCreateUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "CREATE TABLE users (id SERIAL PRIMARY KEY, name TEXT);")
	_, err = tx.ExecContext(ctx, "CREATE TABLE company (id SERIAL PRIMARY KEY, name TEXT);")
	return err
}

func DownCreateUsers(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE users;")
	_, err = tx.ExecContext(ctx, "DROP TABLE company;")
	return err
}
