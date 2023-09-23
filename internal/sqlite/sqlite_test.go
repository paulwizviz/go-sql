package sqlite

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestPrepareStmt(t *testing.T) {

	cases := []struct {
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
			input:       "CREATE TABL IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
			want:        ErrStmt,
			description: "Invalid TABLE token",
		},
	}

	for i, c := range cases {
		db, err := ConnectMem(0, 0, 0, 0)
		if err != nil {
			t.Fatalf(err.Error())
		}
		got := CreateTable(db, c.input)
		if c.want == nil {
			if c.want != got {
				t.Fatalf("Case: %d Description: %s Got: %v Want: %v", i, c.description, got, c.want)
			}
		} else if !errors.Is(got, c.want) {
			t.Fatalf("Case: %d Description: %s Got: %v Want: %v", i, c.description, got, c.want)
		}
	}

}

func BenchmarkConn(b *testing.B) {
	for idx, c := range []Config{
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     0,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(1 * time.Second),
			ConnMaxLifeTime: time.Duration(1 * time.Second),
			MaxIdleConn:     3,
			MaxOpenConn:     3,
		},
	} {
		b.Run(fmt.Sprintf("Scenario %d", idx), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConnectMem(c.ConnMaxIdleTime, c.ConnMaxLifeTime, c.MaxIdleConn, c.MaxOpenConn)
			}
		})
	}
}

func BenchmarkCreateTable(b *testing.B) {
	for idx, c := range []Config{
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     0,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(1 * time.Second),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     0,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(1 * time.Second),
			MaxIdleConn:     0,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     4,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(1 * time.Second),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     4,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     0,
			MaxOpenConn:     4,
		},
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(1 * time.Second),
			MaxIdleConn:     0,
			MaxOpenConn:     4,
		},
	} {
		db, _ := ConnectMem(c.ConnMaxIdleTime, c.ConnMaxLifeTime, c.MaxIdleConn, c.MaxOpenConn)
		b.Run(fmt.Sprintf("Scenario %d", idx), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var strStmts = []string{
					"CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
					"CREATE TABLE IF NOT EXISTS human (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
					"CREATE TABLE IF NOT EXISTS girls (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
					"CREATE TABLE IF NOT EXISTS boys (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)",
					"CREATE TABLE IF NOT EXISTS animals (id INTEGER PRIMARY KEY, name TEXT)",
					"CREATE TABLE IF NOT EXISTS dogs (id INTEGER PRIMARY KEY, name TEXT)",
					"CREATE TABLE IF NOT EXISTS cats (id INTEGER PRIMARY KEY, name TEXT)",
				}
				for _, strStmt := range strStmts {
					CreateTable(db, strStmt)
				}
			}
		})
	}
}
