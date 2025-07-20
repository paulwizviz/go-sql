package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql/internal/person"
	"go-sql/internal/sqlops"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func insertPersonDetail(ctx context.Context, db *sql.DB, details ...person.Detail) error {

	tx, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	insertPersonStmt, err := tx.PrepareContext(ctx, person.SQLiteInsertPersonSQL)
	if err != nil {
		return err
	}
	defer insertPersonStmt.Close()

	insertNameIDStmt, err := tx.PrepareContext(ctx, person.SQLiteInsertNamedIDSQL)
	if err != nil {
		return err
	}
	defer insertNameIDStmt.Close()

	insertPersonNameIDStmt, err := tx.PrepareContext(ctx, person.SQLiteInsertPersonNameIDSQL)
	if err != nil {
		return err
	}
	defer insertPersonNameIDStmt.Close()

	for _, detail := range details {
		result, err := insertPersonStmt.ExecContext(ctx)
		if err != nil {
			return err
		}
		personID, err := result.LastInsertId()
		if err != nil {
			return err
		}
		result, err = insertNameIDStmt.ExecContext(ctx, detail.FirstName, detail.Surname, detail.Nickname)
		if err != nil {
			return err
		}
		nameID, err := result.LastInsertId()
		if err != nil {
			return err
		}
		_, err = insertPersonNameIDStmt.ExecContext(ctx, personID, nameID)
		if err != nil {
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetPersonDetail(ctx context.Context, db *sql.DB) error {
	stmt, err := db.PrepareContext(ctx, person.SelectPersonSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	var (
		personID  int
		nameID    int
		firstName string
		surname   string
		nickname  string
	)
	for rows.Next() {
		err := rows.Scan(
			&personID,
			&nameID,
			&firstName,
			&surname,
			&nickname,
		)
		if err != nil {
			log.Printf("%s", err.Error())
		}
		fmt.Println(personID, nameID, firstName, surname, nickname)
	}

	return nil
}

func main() {
	db, err := sqlops.NewSQLiteMem()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.TODO()

	log.Println("STEP 1: Create Tables")
	err = sqlops.CreateTableTx(
		ctx,
		db,
		person.SQLiteCreatePersonTxFn,
		person.SQLiteCreateNamedIDTxFn,
		person.SQLiteCreatePersonNameIDTxFn,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STEP 2: Insert Person Detail")
	detail := person.Detail{
		FirstName: "John",
		Surname:   "Doe",
		Nickname:  "",
	}

	err = insertPersonDetail(ctx, db, detail)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STEP 3: Get Person Detail")
	err = GetPersonDetail(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

}
