# SQLite

This section discusses all things related to SQLite.

## Useful references

* [Offical Documentation](https://www.sqlite.org/index.html)
* [Connection Strings](https://www.connectionstrings.com/sqlite/)

## SQLite and Go

The Go SQL interface packages do not implement the operations of the underlying databases. To learn more about coding with SQL, please refer to the "[An Introduction to using SQL Databases in Go](https://www.alexedwards.net/blog/introduction-to-using-sql-databases-in-go)"

When you write Go code interacting with SQL databases you must import two packages:

```
	"database/sql" // SQL interfaces
	_ "github.com/mattn/go-sqlite3" // Drivers for SQLite -- e.g. SQLite
```

Please refer to [Golang SQLite database/sql](https://earthly.dev/blog/golang-sqlite/) for techniques to get Go to work with SQLite

