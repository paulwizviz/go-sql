package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"testing"
	"time"
)

type config struct {
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
}

func defaulSQLiteFile() string {
	pwd, _ := os.Getwd()
	dbPath := path.Join(pwd, "tmp", "sqlite")
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	return path.Join(pwd, "tmp", "sqlite", "data.db")
}

func BenchmarkCreateTable(b *testing.B) {

	scenarios := []struct {
		db     func() *sql.DB
		config config
		title  string
	}{
		{
			db: func() *sql.DB {
				db, _ := ConnectMemDefault()
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite default memory",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemDefault()
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(1 * time.Second),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite default memory",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemDefault()
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(1 * time.Second),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite default memory",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemDefault()
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    4,
				MaxOpenConns:    0,
			},
			title: "SQLite default memory",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemDefault()
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    4,
			},
			title: "SQLite default memory",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemPool(defaulSQLiteFile(), 100)
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite memory 100 pool",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemPool(defaulSQLiteFile(), 100)
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(1 * time.Second),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite memory 100 pool",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemPool(defaulSQLiteFile(), 100)
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(1 * time.Second),
				MaxIdleConns:    0,
				MaxOpenConns:    0,
			},
			title: "SQLite memory 100 pool",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemPool(defaulSQLiteFile(), 100)
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    4,
				MaxOpenConns:    0,
			},
			title: "SQLite memory 100 pool",
		},
		{
			db: func() *sql.DB {
				db, _ := ConnectMemPool(defaulSQLiteFile(), 100)
				return db
			},
			config: config{
				ConnMaxIdleTime: time.Duration(0),
				ConnMaxLifeTime: time.Duration(0),
				MaxIdleConns:    0,
				MaxOpenConns:    4,
			},
			title: "SQLite memory 100 pool",
		},
	}

	for idx, s := range scenarios {
		db := s.db()
		defer db.Close()
		b.Run(fmt.Sprintf("Scenario %s %d", s.title, idx), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				db.SetConnMaxIdleTime(s.config.ConnMaxIdleTime)
				db.SetConnMaxLifetime(s.config.ConnMaxLifeTime)
				db.SetMaxIdleConns(s.config.MaxIdleConns)
				db.SetMaxOpenConns(s.config.MaxOpenConns)
				db.Exec("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
				db.Exec("CREATE TABLE IF NOT EXISTS humans (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
				db.Exec("CREATE TABLE IF NOT EXISTS boys (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
				db.Exec("CREATE TABLE IF NOT EXISTS girls (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
			}
		})
	}

}
