# PostgreSQL

This section discuss patterns for implementing Go interaction with PostgreSQL DB.

## Drivers

### lib/pq

This is a older and widely used package. The full package path `github.com/lib/pq`. This is effectively unmaintained — its own `README.md` recommends migrating to `pgx`.

### pgx

This is more modern, feature-rich, often preferred now. The full path is `github.com/jackc/pgx/v5`.

#### Batch & Pipeline Mode

`pgx` lets you send multiple queries in a single round trip:

```go
// Batch - multiple queries, one round trip
batch := &pgx.Batch{}
batch.Queue("INSERT INTO events VALUES ($1)", e1)
batch.Queue("INSERT INTO events VALUES ($1)", e2)
batch.Queue("INSERT INTO events VALUES ($1)", e3)

results := pool.SendBatch(ctx, batch)
defer results.Close()
```

This is huge for write-heavy workloads where network latency dominates.

#### Richer PostgreSQL Type Support

`pgx` natively understands Postgres types that database/sql has no concept of:

```go
// Arrays
var tags []string
row.Scan(&tags) // []text[] just works

// hstore, jsonb, uuid, ranges, composites
// all handled natively without manual marshaling
```

With `database/sql` you're often converting everything to/from strings manually.

#### Built-in Connection Pooling - `pgxpool`

`pgx` ships its own pool tuned specifically for Postgres behaviour — health checks, reconnection, min/max idle connections. `database/sql` has a generic pool that works, but gives you less control and visibility.

#### Context & Cancellation

`pgx` was designed from the ground up with `context.Context` — cancellation propagates cleanly to the Postgres server, releasing server-side resources immediately. With older drivers this was bolted on and often incomplete.

### Comparison

| Feature | `pgx` | `lib/pq` |
| --- | --- | --- |
| Binary protocol | ✅ | ❌ (text only) |
| Server-side prepared statements | ✅ automatic | ❌ |
| Batch queries | ✅ | ❌ |
| Pipeline mode | ✅ | ❌ |

- **Binary protocol** is the big one — instead of sending "12345" as a string over the wire, it sends the actual integer bytes. Faster serialization, less parsing overhead on both ends.

- **Prepared statements** are cached automatically — the query plan is compiled once on the server and reused, which matters at high query volume.

## Working Examples

These examples uses [learn sql](https://github.com/paulwizviz/learn-sql) local deployment as deployed example.

### Example 1: Ping

This [example](../examples/coding/pg/ex1/main.go) involves establishing a connection with a PostgreSQL server and followed by a ping. In this case, the package `lib/pq` is used.

### Example 2: Simple CRUD using lib/pq

This [example](../examples/coding/pg/ex2/main.go) illustrate the process to create table, inserting data, querying and dropping table. In this case, the package `lib/pq` is used.

### Example 3: Simple CRUD using pgx

This [example](../examples/coding/pg/ex3/main.go) is a reproduction of Example 2, but using the package `pgx`.
