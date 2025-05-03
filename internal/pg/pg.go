package pg

import (
	"database/sql"
	"fmt"
)

func NewDB(username string, password string, host string, port uint, dbname string) (*sql.DB, error) {
	connStmt := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := sql.Open("postgres", connStmt)
	if err != nil {
		return nil, err
	}
	return db, nil
}
