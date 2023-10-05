# SQLite

This section discusses all things related to SQLite.

## Useful references

* [Offical Documentation](https://www.sqlite.org/index.html)
* [Connection Strings](https://www.connectionstrings.com/sqlite/)
* [SQLite for Beginners](https://www.youtube.com/watch?v=Wd5WWVx3aRE&list=PLWENznQwkAoxww-cDEfIJ-uuPDfFwbeiJ)
* [Building SQLite with CGo for (almost) every OS](https://zig.news/kristoff/building-sqlite-with-cgo-for-every-os-4cic)

## Data types

* `null` - Includes any NULL values.
* `integer` - Signed integers, stored in 1, 2, 3, 4, 6, or 8 bytes depending on the magnitude of the value.
* `real` - Real numbers, or floating point values, stored as 8-byte floating point numbers.
* `text` - Text strings stored using the database encoding, which can either be UTF-8, UTF-16BE or UTF-16LE.
* `blob` - Any blob of data, with every blob stored exactly as it was input.

## SQLite and Go

The Go SQL interface packages do not implement the operations of the underlying databases. To learn more about coding with SQL, please refer to the "[An Introduction to using SQL Databases in Go](https://www.alexedwards.net/blog/introduction-to-using-sql-databases-in-go)"

When you write Go code interacting with SQL databases you must import two packages:

```
	"database/sql" // SQL interfaces
	_ "github.com/mattn/go-sqlite3" // Drivers for SQLite -- e.g. SQLite
```

Please refer to [Golang SQLite database/sql](https://earthly.dev/blog/golang-sqlite/) for techniques to get Go to work with SQLite

