package sqlops_test

import (
	"context"
	"database/sql"
	"fmt"
	"go-sql/internal/sqlops"
	"reflect"
)

var (
	sqlitePerson = `CREATE TABLE IF NOT EXISTS person (
first_name TEXT,
surname TEXT,
age INTEGER)`

	pgPerson = `CREATE TABLE IF NOT EXISTS person (
first_name VARCHAR(10),
surname VARCHAR(10),
age INT)`

	createPersonTblFunc sqlops.TblCreatorFunc = func(ctx context.Context, db *sql.DB) error {
		dbType := sqlops.DriverType(db)
		switch dbType {
		case sqlops.SQLiteType:
			_, err := db.ExecContext(ctx, sqlitePerson)
			if err != nil {
				return fmt.Errorf("%w-person table", sqlops.ErrTable)
			}
		case sqlops.PSQLType:
			_, err := db.ExecContext(ctx, pgPerson)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("%w-driver %s unsupported", sqlops.ErrTable, reflect.TypeOf(db).String())
		}
		return nil
	}
)

func Example_createTable() {
	db, _ := sqlops.NewSQLiteMem()
	err := sqlops.CreateTable(context.TODO(), createPersonTblFunc, db)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output:
}
