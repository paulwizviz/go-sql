package person

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"go-sql/internal/sqlops"
)

var (

	// Errors
	ErrTblPerson       = errors.New("unable to create person table")
	ErrTblNamedID      = errors.New("unable to create named_identifier table")
	ErrTblPersonNameID = errors.New("unable to create person_name_identifier table")

	//go:embed sql/select_person.sql
	SelectPersonSQL string
)

// NameValue is a type of object value
type NameValue string

// NameIdentifier represent a name based identifier
type NameIdentifier struct {
	PersonID  int       `json:"person_id"`
	NameID    int       `json:"name_id"`
	FirstName NameValue `json:"first_name"`
	Surname   NameValue `json:"middlename"`
	Nickname  NameValue `json:"nickname"`
}

func GetNames(ctx context.Context, db *sql.DB) (NameIdentifier, error) {

	ni := NameIdentifier{}

	var getPersonFn sqlops.RowProcessor = func(rows *sql.Rows) error {
		var personID, nameID int
		var firstName, surname, nickname string
		for rows.Next() {
			err := rows.Scan(&personID, &nameID, &firstName, &surname, &nickname)
			if err != nil {
				return fmt.Errorf("unable to extract name: %v", err)
			}
		}
		ni.PersonID = personID
		ni.NameID = nameID
		ni.FirstName = NameValue(firstName)
		ni.Nickname = NameValue(nickname)
		ni.Surname = NameValue(surname)
		return nil
	}

	sr, err := sqlops.NewStmtReader(ctx, db, SelectPersonSQL)
	if err != nil {
		return NameIdentifier{}, err
	}

	err = sr.Query(ctx, getPersonFn)
	if err != nil {
		return NameIdentifier{}, err
	}
	return ni, nil
}
