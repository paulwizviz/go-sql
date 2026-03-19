# SQLite

This section discuss the techniques for creating application that connect to SQLite database.

## Drivers

### go-sqlite3

This is a CGO based driver. The full package path `github.com/mattn/go-sqlite3`.

### smodernc.org/sqlite

This is a CGO free solution. The full package path is `modernc.org/sqlite`.

## Working Examples

### Example 1: Simple CRUD for SQLite

This [example](../examples/coding/sqlite/ex1/main.go) illustrate the process to instantiate a SQLite server, create table, inserting data, querying and dropping table.

### Example 2: This demonstrates transactions

This [example](../examples/coding/sqlite/ex2/main.go) illustrates a table with only Primary key and the process uses transactions to insert data commit.

### Example 3: Mapping SQLite data to Go custom type

This example is based on this logical schema.

![img person-relation](../assets/img/person-name.png).

The SQL specification of the schema are based on these files:

* [Person Table](../internal/person/sql/sqlite/tbl_person.sql)
* [Name ID Table](../internal/person/sql/sqlite/tbl_named_id.sql)
* [Person Name ID Table](../internal/person/sql/sqlite/tbl_person_name_id.sql)

Here is [the implementation](../examples/coding/sqlite/ex3/main.go)
