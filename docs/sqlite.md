# SQLite

SQLite is a C-language library that implements a small, fast, self-contained, high-reliability, full-featured, SQL database engine. SQLite is the most used database engine in the world. SQLite is built into all mobile phones and most computers and comes bundled inside countless other applications that people use every day ([Source: Official Doc](https://www.sqlite.org/about.html)).

## SQLite and Go

The Go SQL interface packages do not implement the operations of the underlying databases. To learn more about coding with SQL, please refer to the "[An Introduction to using SQL Databases in Go](https://www.alexedwards.net/blog/introduction-to-using-sql-databases-in-go)"

When you write Go code interacting with SQL databases you must import two packages:

```
	"database/sql" // SQL interfaces
	_ "github.com/mattn/go-sqlite3" // Drivers for SQLite -- e.g. SQLite
```

Please refer to [Golang SQLite database/sql](https://earthly.dev/blog/golang-sqlite/) for techniques to get Go to work with SQLite

## A simple example

To help you appreciate the basic operations of Go interacting with SQLite, please refer to this [source code](../cmd/sqlitecmd/cli/main.go). To ensure that it can run in your platform, the example is built and executable from a Docker container. Scripts have been provided to make it easy for you to build and run the app. The script is found [here](../scripts/sqlitedb/cli.sh).

* Build (`./scripts/sqlitedb/cli build`) to build the cli app
* Clean (`./scripts/sqlitedb/cli clean`) to remove traces of cli app
* Run (`./scripts/sqlitedb/cli shell`) to open a shell in a container. Then run `sqlitecmd`