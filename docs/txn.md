# Transaction

An SQL transaction is a single logical unit of work that comprises one or more SQL statements. The key characteristic of a transaction is that it is treated as an atomic operation. This means either all the statements within the transaction are successfully executed and committed to the database, or none of them are. If any part of the transaction fails, the entire transaction is rolled back, and the database is returned to its state before the transaction began.

The properties that define a transaction are often referred to by the acronym ACID:

* **Atomicity**: As explained above, a transaction is an indivisible unit. It either completes entirely or has no effect.
* **Consistency**: A transaction brings the database from one valid state to another. It ensures that data conforms to all defined rules and constraints (e.g., primary key constraints, foreign key constraints, check constraints).
* **Isolation**: The execution of concurrent transactions should produce the same result as if they were executed sequentially. This means that one transaction's partial or uncommitted changes should not be visible to other transactions. Different isolation levels exist (e.g., Read Uncommitted, Read Committed, Repeatable Read, Serializable) to control the degree of isolation.
* **Durability**: Once a transaction has been committed, its changes are permanent and will survive system failures (e.g., power outages, crashes). The committed data is stored persistently.

WHy are Transactions Important?

Consider a bank transfer from Account A to Account B. This involves two operations:

* **STEP 1**: Decrementing the balance of Account A.
* **STEP 2**: Incrementing the balance of Account B.

If these two operations are not treated as a single transaction, and the system crashes after step 1 but before step 2, Account A would be debited, but Account B would not be credited. This would lead to a loss of money and an inconsistent database state. By wrapping these operations in a transaction, either both succeed, or neither does, ensuring the financial integrity.

## Using Go `database/sql` Package

There are two ways to initiate a database transaction:

* `db.Begin()`
* `db.BeginTx(ctx context.Context, opts *sql.TxOptions)`

### `db.Begin()`

* **Signature**: `func (db *DB) Begin() (*Tx, error)`
* **Purpose**: This is the simpler of the two. It starts a new transaction using a default context (`context.Background()`) and the database's default isolation level.
* **Use Case**: Ideal for straightforward transactions where you don't need to specify custom options like a specific context, isolation level, or read-only mode. It's often used when you're sure the default behavior is sufficient and you don't have a `context.Context` readily available for propagation.

### `db.BeginTx(ctx context.Context, opts *sql.TxOptions)`

* **Signature**: `func (db *DB) BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)`
* **Purpose**: This method provides more granular control over the transaction's behavior by allowing you to specify a context.Context and *sql.TxOptions.
* **Use Cases**:
  * **Context Propagation**: The most significant advantage is the ability to pass a `context.Context`. This context is crucial for:
    * **Timeouts and Deadlines**: You can set a timeout for the entire transaction (e.g., using `context.WithTimeout`). If the context is canceled before the transaction is committed or rolled back, the `sql` package will automatically roll back the transaction. This is vital for preventing long-running or stalled transactions from holding database connections indefinitely.
    * **Cancellation Signals**: If an upstream operation (e.g., an HTTP request) is canceled, you can propagate that cancellation to the database transaction, causing it to roll back and free up resources.
    * **Request-Scoped Values**: While less common for transactions directly, contexts can carry request-scoped values, though this is usually applied to individual queries within the transaction rather than the transaction itself.
  * **Isolation Level**: You can explicitly set the transaction's isolation level using `opts.Isolation`. This allows you to choose between different levels like `sql.LevelSerializable`, `sql.LevelRepeatableRead`, `sql.LevelReadCommitted`, or `sql.LevelReadUncommitted`, depending on your application's concurrency and consistency requirements.
  * **Read-Only**: You can set opts.ReadOnly = true to indicate that the transaction will only perform read operations. This can sometimes allow the database to optimize resource usage or prevent accidental writes.

### `sql.TxOptions` strut

The sql.TxOptions struct has the following fields:

```go
type TxOptions struct {
    Isolation IsolationLevel
    ReadOnly  bool
}
```

* `Isolation IsolationLevel`: This field specifies the isolation level for the transaction. `IsolationLevel` is an enum with values like:
  * `sql.LevelDefault`: Uses the database's default isolation level.
  * `sql.LevelReadUncommitted`: Allows "dirty reads" (reading uncommitted changes from other transactions).
  * `sql.LevelReadCommitted`: Prevents dirty reads but allows "non-repeatable reads" (reading the same row twice within a transaction yields different results if another committed transaction changed it). This is often the default for many databases.
  * `sql.LevelRepeatableRead`: Prevents dirty reads and non-repeatable reads but can suffer from "phantom reads" (new rows appearing in result sets of queries within the same transaction).
  * `sql.LevelSerializable`: The highest isolation level. It prevents all the above phenomena, ensuring that concurrent transactions execute as if they were serialized. This offers the strongest consistency but can significantly impact performance due to increased locking.
* `ReadOnly bool`: If set to `true`, indicates that the transaction will only read data and not modify it. This can allow for database optimizations or prevent accidental writes in a read-only context.

## SQL Operations Suited for Transactions

The `database/sql` package's support for transactions means it provides the API to initiate, commit, and rollback transactions, regardless of the underlying database. Any relational database for which there is a compliant Go `database/sql` driver will support transactions through the `database/sql` package.

Here's a breakdown of the types of SQL operations and situations where transactions are essential.

### Data Modification Language (DML) Operations

Transactions are primarily used with DML statements:

* **INSERT**: When you need to insert related data into multiple tables.
  * **Example**: Creating a new order in an e-commerce system. This might involve:
    * `INSERT` into an `orders` table.
    * `INSERT` into an `order_items` table for each product.
    * `UPDATE` a products table to decrement inventory.
    * If any of these fail (e.g., insufficient stock), the entire order creation should be rolled back.
* **UPDATE**: When updating multiple interdependent rows or tables.
  * **Example**: Transferring money between bank accounts:
    * `UPDATE` Account A (debit).
    * `UPDATE` Account B (credit).
    * These two updates must both succeed or both fail.
* `DELETE`: When deleting data that has referential integrity constraints or requires cascading deletions.
  * **Example**: Deleting a customer:
    * `DELETE` from the `customers` table.
    * `DELETE` related records from `orders`, `addresses`, `payments`, etc. (if not handled by foreign key cascading rules).
    * If the customer is deleted but their orders remain, it's an inconsistent state.

### Operations Requiring Atomicity (All or Nothing)

Any scenario where a set of operations must either entirely succeed or entirely fail. This is the definition of atomicity, the "A" in ACID.

* **Financial Transactions**: As in the bank transfer example, any movement of funds, stock trades, or accounting entries.
* **Inventory Management**: Decrementing stock when an item is sold, and then creating an order record.
* **User Registration**: Creating a user record, a profile record, and perhaps default settings records.
* **Workflow Steps**: If a business process involves multiple database changes, and the whole process must be completed, or none of it should be.

### Maintaining Consistency (ACID "C")

Transactions help enforce business rules and constraints. If a set of changes would temporarily violate a constraint (e.g., an intermediate step makes a balance negative before another step corrects it), a transaction ensures that the database only sees the final, consistent state.

### Isolation (ACID "I")

When multiple users or processes are concurrently accessing and modifying the same data, transactions provide isolation. This prevents one transaction's uncommitted changes from interfering with another, avoiding issues like:

* **Dirty Reads**: Reading data that has been modified by another transaction but not yet committed.
* **Non-Repeatable Reads**: Reading the same row multiple times within a single transaction yields different values because another transaction committed a change to that row in between.
* **Phantom Reads**: A query within a transaction returns a set of rows, and later the same query returns more rows because another transaction inserted new rows that match the query's criteria.

While a single SELECT statement doesn't usually require a transaction (as it doesn't modify data), read operations within a transaction are crucial for maintaining consistency and isolation. For example, if you need to read a set of data and then update it based on that read, you'd want both the read and the update to be within the same transaction to prevent another process from changing the data between your read and write.

### Error Recovery

Transactions provide a clear mechanism for error recovery. If an error occurs at any point during a multi-step operation, a ROLLBACK ensures that the database reverts to its state before the transaction began, preventing partial updates and data corruption.

When NOT to use Transactions (or be cautious):

* **Long-Running Operations**: Transactions acquire locks on data, which can block other operations. Very long transactions reduce concurrency and can lead to performance bottlenecks, deadlocks, and connection pool exhaustion.
* **Batch Processing (sometimes)**: For extremely large batch operations, a single transaction might be too large. It's often better to break them into smaller, manageable transactions to reduce lock contention and allow for partial progress.
* **Read-Only Operations (standalone SELECT)**: A simple SELECT query that doesn't need to be consistent with subsequent writes (or guarantee isolation from other concurrent writes) typically doesn't need to be wrapped in an explicit transaction. Many databases implicitly wrap single DML statements in an "auto-commit" transaction anyway.

In essence, if your database operation involves multiple, interdependent steps that must all succeed or all fail to maintain data integrity, it's a prime candidate for an SQL transaction.

### Data Definition Language (DDL)

DDL (Data Definition Language) operations (like CREATE TABLE, ALTER TABLE, DROP TABLE, CREATE INDEX, DROP INDEX) are different. They modify the schema (structure) of the database, not just the data within it. The transactional behavior of DDL is database-specific.

Here's how popular databases typically handle DDL within transactions:

* **PostgreSQL: Full Transactional DDL (Mostly Yes!)**
  * Yes, PostgreSQL is a standout here. It supports transactional DDL very well. You can include `CREATE TABLE`, `ALTER TABLE`, `DROP TABLE`, `CREATE INDEX`, etc., within a `BEGIN; ... COMMIT;` block.
  * If an error occurs, or you explicitly ROLLBACK;, all DDL changes made within that transaction will be undone, and the database schema will revert to its state before the transaction began.
  * **Exceptions**: Some truly global operations, like CREATE DATABASE or DROP TABLESPACE, might not be fully reversible, but for typical schema changes, it's transactional.
  * **Benefit**: This is incredibly powerful for schema migrations. You can script a complex series of schema changes, and if any step fails, you can roll back the entire migration, preventing a partially updated or broken schema.
* **MySQL: Limited / Implicit Commit (Generally No for Multi-Statement Rollback)**
  * No, traditionally MySQL (especially with InnoDB, the transactional engine) does not support transactional DDL in the way PostgreSQL does.
  * Implicit Commit: Most DDL statements in MySQL (e.g., `CREATE TABLE`, `ALTER TABLE`, `DROP TABLE`) cause an implicit commit of any currently active transaction before the DDL statement executes, and often after it executes as well. This means you cannot reliably roll back DDL operations.
  * **Atomic DDL (MySQL 8.0+)**: MySQL 8.0 introduced "Atomic DDL." This means that individual DDL statements are atomic (they either fully succeed or fully fail). If a `CREATE TABLE` statement fails, it won't leave behind a partial table. However, this does not mean you can put multiple DDL statements in a transaction and roll them all back if one fails. Each DDL statement still implicitly commits.
  * **Implication**: For MySQL, you typically manage schema changes carefully outside of a single transaction, or you rely on migration tools that manage schema versions and rollback strategies externally (e.g., by applying "down" migrations).
* **SQL Server: Partial Support**
  * SQL Server offers some level of transactional DDL support, particularly with SAVEPOINTs, but it can be more nuanced than PostgreSQL.
  * You can often enclose DDL in BEGIN TRANSACTION; ... COMMIT; blocks, and many DDL operations will roll back.
  * However, there are exceptions, and the behavior can sometimes be unexpected for certain DDL commands or complex scenarios. It's generally safer to test thoroughly or rely on specific tools for schema management.
* **Oracle: No (Implicit Commit)**
  * Similar to MySQL, Oracle typically performs an implicit commit before and after DDL statements. This means DDL statements cannot be rolled back in the traditional sense within a transaction.
  * Oracle has other mechanisms for schema management and versioning (like flashback features), but not standard transactional DDL.
* SQLite: Yes (Full Transactional DDL)
  * Yes, SQLite fully supports transactional DDL. You can include CREATE TABLE, ALTER TABLE, etc., inside a transaction, and they will be rolled back if the transaction is rolled back. This is one of SQLite's strengths for simpler, file-based database management.
