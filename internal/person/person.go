package person

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"go-sql/internal/sqlops"
	"log"
	"reflect"
)

var (
	ErrTable           = errors.New("undefined table")
	ErrTblPerson       = errors.New("person")
	ErrTblNamedID      = errors.New("named_identifier")
	ErrTblPersonNameID = errors.New("person_name_identifier")
)

var (
	//go:embed sql/sqlite/tbl_person.sql
	sqllitePersonSQL string
	//go:embed sql/sqlite/tbl_named_id.sql
	sqliteNameIDSQL string
	//go:embed sql/sqlite/tbl_person_name_id.sql
	sqlitePersonNameIDSQL string

	createPersonTblFunc sqlops.TblCreatorFunc = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, sqllitePersonSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPerson, err)
		}
		return nil
	}

	createNamedIDTblFunc sqlops.TblCreatorFunc = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, sqliteNameIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblNamedID, err)
		}
		return nil
	}

	createPersonNameIDTblFunc sqlops.TblCreatorFunc = func(ctx context.Context, db *sql.DB) error {
		_, err := db.ExecContext(ctx, sqlitePersonNameIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPersonNameID, err)
		}
		return nil
	}
)

func CreateTables(ctx context.Context, db *sql.DB) {
	dbType := sqlops.DriverType(db)
	switch dbType {
	case sqlops.SQLiteType:
		if err := sqlops.CreateTable(ctx, createPersonTblFunc, db); err != nil {
			log.Println(err)
		}
		if err := sqlops.CreateTable(ctx, createNamedIDTblFunc, db); err != nil {
			log.Println(err)
		}
		if err := sqlops.CreateTable(ctx, createPersonNameIDTblFunc, db); err != nil {
			log.Println(err)
		}
	default:
		log.Println(fmt.Errorf("%w-driver %s unsupported", ErrTable, reflect.TypeOf(db).String()))
	}
}
