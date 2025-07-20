package person

import (
	_ "embed"
	"errors"
)

var (

	// Errors
	ErrTblPerson       = errors.New("unable to create person table")
	ErrTblNamedID      = errors.New("unable to create named_identifier table")
	ErrTblPersonNameID = errors.New("unable to create person_name_identifier table")

	//go:embed sql/select_person.sql
	SelectPersonSQL string
)

type Detail struct {
	FirstName string `json:"first_name"`
	Surname   string `json:"surname"`
	Nickname  string `json:"nick_name"`
}
