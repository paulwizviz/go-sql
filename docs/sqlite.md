# SQLite

[Offical Documentation](https://www.sqlite.org/index.html)

## SQLite and Go

The Go SQL interface packages do not implement the operations of the underlying databases. To learn more about coding with SQL, please refer to the "[An Introduction to using SQL Databases in Go](https://www.alexedwards.net/blog/introduction-to-using-sql-databases-in-go)"

When you write Go code interacting with SQL databases you must import two packages:

```
	"database/sql" // SQL interfaces
	_ "github.com/mattn/go-sqlite3" // Drivers for SQLite -- e.g. SQLite
```

Please refer to [Golang SQLite database/sql](https://earthly.dev/blog/golang-sqlite/) for techniques to get Go to work with SQLite

## Working Examples

A series of examples are built into an [Ubuntu container](../build/sqlite/sqlite.dockerfile)

### Ex1

This example demonstrate simple application to store in SQLite file. [Source](../cmd/sqlite/ex1/main.go)

### Ex2

This example demonstrates primary key constrains error.