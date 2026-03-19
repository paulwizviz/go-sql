package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	createTableStmtStr = "CREATE TABLE IF NOT EXISTS lottery(id SERIAL PRIMARY KEY, ball1 INT, ball2 INT)"
	insertStmtStr      = "INSERT INTO lottery (ball1, ball2) VALUES ($1,$2) RETURNING id"
	selectStmtStr      = "SELECT ball1, ball2 FROM lottery WHERE ball1=$1 AND ball2=$2"
	dropTableStmtStr   = "DROP TABLE lottery"
)

func insertStatement(pool *pgxpool.Pool, ctx context.Context, args []int) error {
	var id int
	err := pool.QueryRow(ctx, insertStmtStr, args[0], args[1]).Scan(&id)
	if err != nil {
		return err
	}
	log.Printf("Inserted row with ID: %v", id)
	return nil
}

func selectQuery(pool *pgxpool.Pool, ctx context.Context, arg1, arg2 int) error {
	rows, err := pool.Query(ctx, selectStmtStr, arg1, arg2)
	if err != nil {
		return err
	}
	defer rows.Close()

	var ball1, ball2 int
	for rows.Next() {
		if err := rows.Scan(&ball1, &ball2); err != nil {
			return err
		}
		fmt.Println(ball1, ball2)
	}
	return rows.Err()
}

func main() {
	ctx := context.Background()

	// STEP 1: Establish connection pool
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/default")
	if err != nil {
		log.Fatal(err)
	}

	// STEP 6: Close pool
	defer pool.Close()

	// STEP 2: Create Table
	if _, err := pool.Exec(ctx, createTableStmtStr); err != nil {
		log.Fatalf("Create table error: %v", err)
	}

	// STEP 3: Insert
	if err := insertStatement(pool, ctx, []int{1, 2}); err != nil {
		log.Fatalf("Insert execution error: %v", err)
	}

	// STEP 4: Select
	if err := selectQuery(pool, ctx, 1, 2); err != nil {
		log.Fatalf("Select query error: %v", err)
	}

	// STEP 5: Drop Table
	if _, err := pool.Exec(ctx, dropTableStmtStr); err != nil {
		log.Fatalf("Drop table error: %v", err)
	}
}
