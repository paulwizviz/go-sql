package person

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"go-sql/internal/sqlops"
)

// SQLite
var (
	//go:embed sql/sqlite/tbl_person.sql
	SQLiteTblPersonSQL string
	//go:embed sql/sqlite/tbl_named_id.sql
	SQLiteTblNamedIDSQL string
	//go:embed sql/sqlite/tbl_person_name_id.sql
	SQLiteTblPersonNameIDSQL string
	//go:embed sql/sqlite/insert_person.sql
	SQLiteInsertPersonSQL string
	//go:embed sql/sqlite/insert_named_id.sql
	SQLiteInsertNamedIDSQL string
	//go:embed sql/sqlite/insert_person_name_id.sql
	SQLiteInsertPersonNameIDSQL string

	SQLiteCreatePersonTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblPersonSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPerson, err)
		}
		return nil
	}

	SQLiteCreateNamedIDTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblNamedIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblNamedID, err)
		}
		return nil
	}

	SQLiteCreatePersonNameIDTxFn sqlops.TblCreatorTxFunc = func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, SQLiteTblPersonNameIDSQL)
		if err != nil {
			return fmt.Errorf("%w-%v", ErrTblPersonNameID, err)
		}
		return nil
	}
)
