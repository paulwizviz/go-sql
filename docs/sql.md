# SQL Syntax

* [Types Differences](#database-type-differences)
* [Statement Parameter Binding](#statement-parameter-binding)

## Database Type Differences

Understanding the differences in SQL concepts and data types across various database systems (SQLite, PostgreSQL, MySQL) is crucial for writing portable and efficient SQL operations. Below is a summary:

| Concept | SQLite | PostgreSQL | MySQL | Notes |
| --- | --- | --- | --- | --- |
| Integer | INTEGER | INTEGER,INT,BIGINT,SMALLINT | INT,BIGINT,TINYINT,SMALLINT,MEDIUMINT  | SQLite's INTEGER is quite flexible and can store various integer sizes. PostgreSQL and MySQL offer more specific integer types for different ranges. |
| Text/String | TEXT | VARCHAR(n),TEXT,CHAR(n) | VARCHAR(n),TEXT,CHAR(n) | `TEXT` in SQLite is typically variable-length. In PostgreSQL and MySQL, `TEXT` is for very long strings, while VARCHAR(n) is for variable-length strings up to n characters. CHAR(n) is fixed-length. |
| Numbers (Decimal/Floating) | REAL, NUMERIC | NUMERIC(p,s),DECIMAL(p,s),REAL,DOUBLE PRECISION | DECIMAL(p,s),NUMERIC(p,s),FLOAT,DOUBLE | REAL in SQLite is a floating-point number. NUMERIC(p,s)/DECIMAL(p,s) are for exact precision (p=precision, s=scale) and are widely supported. FLOAT and DOUBLE are for approximate floating-point numbers. |
| Boolean | INTEGER (0 for false, 1 for true) | BOOLEAN, BOOL | TINYINT(1) (0 for false, 1 for true) | SQLite doesn't have a native boolean type, often using INTEGER instead. MySQL often uses TINYINT(1) for boolean, and PostgreSQL has a dedicated BOOLEAN type. |
| Date/Time | TEXT,INTEGER,REAL | DATE,TIME,TIMESTAMP,TIMESTAMPTZ | DATE,TIME,DATETIME,TIMESTAMP | SQLite stores dates/times as text (ISO8601 strings), integers (Unix epoch time), or real numbers (Julian day numbers). PostgreSQL and MySQL have dedicated and more robust date/time types, including options for time zones (TIMESTAMPTZ in PostgreSQL). |
| Binary Data | BLOB | BYTEA | BLOB,TINYBLOB,MEDIUMBLOB,LONGBLOB | All support binary large objects. |

## Statement Parameter Binding

| Database | Positional Anonymous | Positional Numbered | Named (Native SQL). | Named (Client/Driver specific) |
| --- | --- | --- | --- | --- |
| SQLite. | ? | ?N  | :name, @name, $name. | Yes (often supports all) |
| MySQL | ? | No | No | Yes (common in client libraries) |
| PostgreSQL | No | $N | No | Yes (common in client libraries) |

## Drivers

Here's a list of some of the most common relational databases and their popular Go drivers that fully support the database/sql transaction model:

* **PostgreSQL**:
  * `github.com/lib/pq` (older, but widely used)
  * `github.com/jackc/pgx/v5` (more modern, feature-rich, often preferred now)
* **MySQL / MariaDB**:
  * `github.com/go-sql-driver/mysql`
* **SQLite**:
  * `github.com/mattn/go-sqlite3`
* **Microsoft SQL Server**:
  * `github.com/denisenkom/go-mssqldb`
* **Oracle Database**:
  * `github.com/go-godbc/godbc` (or other commercial/community drivers)
