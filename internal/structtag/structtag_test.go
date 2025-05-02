package structtag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqliteDBTags(t *testing.T) {
	tagname := "sqlite"

	testcases := []struct {
		input       any
		expected    []StructTag
		description string
	}{
		{
			input: &Human{},
			expected: []StructTag{
				{
					FieldName: "ID",
					Tag:       "id,INTEGER,PRIMARY_KEY",
				},
				{
					FieldName: "FirstName",
					Tag:       "first_name,TEXT",
				},
				{
					FieldName: "Surname",
					Tag:       "surname,TEXT",
				},
			},
			description: "Human tags",
		},
		{
			input: &Animal{},
			expected: []StructTag{
				{
					FieldName: "ID",
					Tag:       "id,INTEGER,PRIMARY_KEY",
				},
				{
					FieldName: "Species",
					Tag:       "species,TEXT",
				},
				{
					FieldName: "Name",
					Tag:       "name,TEXT",
				},
			},
			description: "Human tags",
		},
	}

	for i, tc := range testcases {
		switch v := tc.input.(type) {
		case *Human:
			actual := DBTags(tagname, v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		case *Animal:
			actual := DBTags(tagname, v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}

	}
}

func TestSQLiteCreateTblStmtStr(t *testing.T) {
	testcases := []struct {
		input       any
		expected    string
		description string
	}{
		{
			input:       &Human{},
			expected:    "CREATE TABLE IF NOT EXISTS human ( id INTEGER PRIMARY KEY, first_name TEXT, surname TEXT )",
			description: "Create human table",
		},
		{
			input:       &Animal{},
			expected:    "CREATE TABLE IF NOT EXISTS animal ( id INTEGER PRIMARY KEY, species TEXT, name TEXT )",
			description: "Create animal table",
		},
	}

	for i, tc := range testcases {
		switch v := tc.input.(type) {
		case *Human:
			actual := SQLiteCreateTblStmtStr(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		case *Animal:
			actual := SQLiteCreateTblStmtStr(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}

func TestSQLiteInsertStmtStr(t *testing.T) {
	testcases := []struct {
		input       any
		expected    string
		description string
	}{
		{
			input:       &Human{},
			expected:    "INSERT INTO human ( id, first_name, surname) VALUES ( ?, ?, ? )",
			description: "Insert into human table",
		},
		{
			input:       &Animal{},
			expected:    "INSERT INTO animal ( id, species, name) VALUES ( ?, ?, ? )",
			description: "Insert into animal table",
		},
	}

	for i, tc := range testcases {
		switch v := tc.input.(type) {
		case *Human:
			actual := SQLiteInsertStmtStr(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		case *Animal:
			actual := SQLiteInsertStmtStr(v)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Case: %d Description: %s", i, tc.description))
		}
	}
}
