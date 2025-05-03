package sqlite

import (
	"errors"
	"fmt"
	"testing"
)

func createTable(input string) error {
	db, err := NewMemDB()
	if err != nil {
		return fmt.Errorf("%w-%s", ErrConn, err.Error())
	}
	_, err = db.Exec(input)
	if err != nil {
		return fmt.Errorf("%w-%s", ErrTable, err.Error())
	}
	return nil
}

// This test is used to verify CREATE TABLE syntax https://www.sqlite.org/lang_createtable.html
func TestCreateTableStatement(t *testing.T) {
	testcases := []struct {
		input       string
		want        error
		description string
	}{
		{
			input:       "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
			want:        nil,
			description: "Valid statement",
		},
		{
			input:       "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname STRING, lastname TEXT)",
			want:        nil,
			description: "Replace TEXT with STRING",
		},
		{
			input:       "CREATE TABLE IF NOT EXISTS people (INTEGER id, firstname STRING, lastname TEXT)",
			want:        nil,
			description: "Syntax err with type declaration",
		},
		{
			input:       "CREAT TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
			want:        ErrTable,
			description: "Syntax err CREAT instead of CREATE",
		},
		{
			input:       "CREATE TABLE IF NOT EXISTS people (PRIMARY KEY id INTEGER, firstname STRING, lastname TEXT)",
			want:        ErrTable,
			description: "Syntax err with constraint declaration",
		},
	}

	for i, tc := range testcases {
		got := createTable(tc.input)
		if tc.want == nil {
			if tc.want != got {
				t.Fatalf("Case: %d Description: %s Want: %v Got: %v", i, tc.description, tc.want, got)
			}
		} else {
			if !errors.Is(got, ErrTable) {
				t.Fatalf("Case: %d Description: %s Want: %v Got: %v", i, tc.description, tc.want, got)
			}
		}
	}
}
