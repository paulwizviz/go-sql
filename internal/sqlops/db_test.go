package sqlops

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/assert"
)

func TestDriverType(t *testing.T) {
	testcases := []struct {
		name      string
		sqlDriver func() (*sql.DB, error)
		expected  SQLType
	}{
		{
			name: "github.com/mattn/go-sqlite3",
			sqlDriver: func() (*sql.DB, error) {
				return NewSQLiteMem()
			},
			expected: SQLiteType,
		},
		{
			name: "github.com/lib/pq",
			sqlDriver: func() (*sql.DB, error) {
				return NewPGConn("a", "b", "c", 123, "efg")
			},
			expected: PSQLType,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d-%s", i, tc.name), func(t *testing.T) {
			db, _ := tc.sqlDriver()
			actual := DriverType(db)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
