package sqlite

import (
	"fmt"
	"testing"
	"time"
)

type config struct {
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	MaxIdleConn     int
	MaxOpenConn     int
}

func BenchmarkConn(b *testing.B) {
	for idx, c := range []config{
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
				Connect(c.ConnMaxIdleTime, c.ConnMaxLifeTime, c.MaxIdleConn, c.MaxOpenConn)
			}
		})
	}
}

func BenchmarkCreateTable(b *testing.B) {
	for idx, c := range []config{
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     0,
			MaxOpenConn:     0,
		},
		{
			ConnMaxIdleTime: time.Duration(0),
			ConnMaxLifeTime: time.Duration(0),
			MaxIdleConn:     2,
			MaxOpenConn:     0,
		},
	} {
		db, _ := Connect(c.ConnMaxIdleTime, c.ConnMaxLifeTime, c.MaxIdleConn, c.MaxOpenConn)
		b.Run(fmt.Sprintf("Scenario %d", idx), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CreateTable(db)
			}
		})
	}
}
