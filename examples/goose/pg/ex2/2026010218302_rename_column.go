package main

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpRootCol, DownRootCol)
}

func UpRootCol(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "UPDATE users SET username='admin' WHERE username='root';")
	return err
}

func DownRootCol(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "UPDATE users SET username='root' WHERE username='admin';")
	return err
}
