# Coding Patterns

## DB Type

- [PostgreSQL](./postgres.md)
- [SQLite](./sqlite.md)

## When to use Prepare Statement

There is no need to use prepare statement for PostgreSQL and SQLite. The `database/sql` package handles statement preparation automatically behind the scenes when you call `db.Query()` or `db.Exec()` directly.

### How database/sql handles it internally

When you call:

```go
db.Query("SELECT * FROM users WHERE id = ?", 1)
```

`database/sql` internally:

- Prepares the statement
- Executes it
- **Caches it** for reuse on the same connection

So subsequent calls with the same query string reuse the cached plan automatically.

### When manual `db.Prepare()` is still useful

The same rules apply as with any driver — explicit preparation is only worth it in specific cases:

```go
// ✅ Worth it — same statement executed many times in a loop
stmt, err := db.Prepare("INSERT INTO events (name) VALUES (?)")
defer stmt.Close()

for _, event := range events {
    stmt.Exec(event)
}
```

This avoids the overhead of the cache lookup on every iteration. But for typical one-off queries, it adds boilerplate with no real benefit.

### Summary of when to use Prepare Statement

| | PostgreSQL (`pgx`) | PostgreSQL (`lib/pq`) | SQLite |
| --- | --- | --- | --- |
| Manual prepare needed? | ❌ | ❌ | ❌ |
| Auto caches statements? | ✅ `pgx` level | ✅ `database/sql` level | ✅ `database/sql` level |
| Manual prepare ever useful? | Rarely | Loop optimisation | Loop optimisation |

The rule of thumb across all three — **let the driver handle it unless you have a specific reason not to**.

## Managing Connections

Here is the discussion related to [managing connections](https://go.dev/doc/database/manage-connections)
