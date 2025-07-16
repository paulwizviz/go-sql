package sqlops

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

// SQLType represents a variant of a SQL base
// on a DB type
type SQLType int

func (s SQLType) String() string {
	switch s {
	case SQLiteType:
		return "SQLite"
	case PSQLType:
		return "PostgreSQL"
	default:
		return "Unspecified"
	}
}

const (
	UnsupportedType SQLType = iota
	SQLiteType
	PSQLType
)

// DriverType returns the SQL type,
func DriverType(db *sql.DB) SQLType {
	switch reflect.TypeOf(db.Driver()).String() {
	case "*sqlite3.SQLiteDriver":
		return SQLiteType
	case "*pq.Driver":
		return PSQLType
	default:
		return UnsupportedType
	}
}

var (
	ErrConn = errors.New("db connect error")
)

const (
	ver = "sqlite3"
)

// NewSQLiteMem instantiate an SQLite Memory connection
func NewSQLiteMem() (*sql.DB, error) {
	db, err := sql.Open(ver, ":memory:")
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

// NewSQLiteFile instantiate an Sqlite file connection
func NewSQLiteFile(f string) (*sql.DB, error) {
	db, err := sql.Open(ver, f)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}

// NewPGConn instantiate a postgress connection
func NewPGConn(username string, password string, host string, port uint, dbname string) (*sql.DB, error) {
	connStmt := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStmt)
	if err != nil {
		return nil, fmt.Errorf("%w-%v", ErrConn, err)
	}
	return db, nil
}
