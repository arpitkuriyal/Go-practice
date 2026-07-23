# 10. Database Operations: Keep Production Databases Fast and Reliable

This is the production-facing database lesson. Learn SQL, migrations, and repository patterns first.

## What “database operations” means

Writing correct SQL is only half the job. In production, you also need to know whether the database is slow, overloaded, blocked, unavailable, or running out of connections—and how to recover data safely.

```text
HTTP request
    ↓
Go handler/service
    ↓ waits for an available connection
sql.DB connection pool
    ↓ sends SQL
PostgreSQL → executes query → reads disk/cache → waits for locks if needed
    ↓
Go scans rows and returns a response
```

An endpoint can feel slow because of any part of this flow. Do not assume “the database is slow” before measuring where time was spent.

## The signals to watch

Use a dashboard and alerts to watch these signals over time. Prefer percentiles (`p95` and `p99`) over only averages: an average can look healthy while a small but important group of requests is very slow.

| Signal | What it tells you | First question to ask |
| --- | --- | --- |
| Query duration (`p50`, `p95`, `p99`) | How long a database operation takes. | Which operation/query name became slower? |
| HTTP request duration | What users actually experience. | Is time in Go code, pool waiting, or the query? |
| Error rate | Failed queries or unavailable database. | Is it a timeout, connection failure, constraint error, or cancellation? |
| Pool wait count/duration | Go requests waited because no connection was free. | Is the pool too small, or are queries/connections held too long? |
| Open/in-use/idle connections | Current pool pressure. | Is `InUse` close to the maximum for long periods? |
| Database CPU, memory, disk IO | Server resource pressure. | Is the database CPU-bound, memory-starved, or reading from disk? |
| Lock waits/deadlocks | Transactions are blocking each other. | Which transaction holds the lock too long? |
| Replication lag (if replicas exist) | Read replicas may return old data. | Can this endpoint safely read stale data? |

## Measure operations, not raw SQL values

Record a stable operation name and duration:

```text
db.operation=find_user duration=12ms outcome=success
db.operation=create_order duration=820ms outcome=timeout
```

Good metadata includes operation name, duration, outcome, and safe error category. Do **not** log passwords, connection strings, tokens, or raw personal data. Be cautious with full SQL text and arguments: they can expose sensitive data and create high-cardinality metrics.

## Connection pools: the important mental model

`*sql.DB` is not one connection. It is a concurrency-safe **pool** that lends connections to queries and receives them back when a query/transaction finishes.

```text
10 HTTP requests
       ↓
pool allows up to 5 open connections
       ↓
5 queries run; 5 requests wait for a connection
```

More connections are not automatically better. Each active PostgreSQL connection consumes resources. A very large pool can overload the database with simultaneous work; a very small pool makes the application wait unnecessarily.

## Configure and observe the pool

[`operations-example.go`](operations-example.go) includes a starting configuration:

```go
db.SetMaxOpenConns(20)                  // maximum connections this process opens
db.SetMaxIdleConns(10)                  // reusable idle connections
db.SetConnMaxLifetime(30 * time.Minute) // periodically replace old connections
db.SetConnMaxIdleTime(5 * time.Minute)  // close unused connections eventually
```

These are examples, not magic production numbers. Start from the database connection limit, reserve capacity for migrations/admin work/other services, divide the remaining budget across application instances, then tune using metrics.

`db.Stats()` provides useful pool data:

| Field | Meaning |
| --- | --- |
| `OpenConnections` | Connections currently open. |
| `InUse` | Connections executing/assigned to work. |
| `Idle` | Open connections ready for reuse. |
| `WaitCount` | Cumulative number of waits for a connection. |
| `WaitDuration` | Cumulative time spent waiting for a connection. |

`WaitCount` and `WaitDuration` are cumulative. Alert on their **rate/change** during a time interval, not just their total since process startup.

## Timeouts and cancellation

Every database operation should receive a context. In an HTTP handler, start with `r.Context()` so work stops when the client disconnects or the server shuts down.

```go
queryContext, cancel := context.WithTimeout(r.Context(), 2*time.Second)
defer cancel()

row := db.QueryRowContext(queryContext, query, id)
```

Use a timeout that fits the operation. A lookup might need hundreds of milliseconds; a report may need longer. A timeout is a safety boundary, not a substitute for fixing a consistently slow query.

At application startup, check connectivity with `db.PingContext` using a bounded context. Do not ping the database for every request.

## When a query becomes slow

1. Identify the operation name and the time it became slower.
2. Check whether requests are waiting for a pool connection or the query itself is slow.
3. Check database CPU, memory, disk IO, connection count, and lock waits.
4. Run the real query with representative values using `EXPLAIN ANALYZE` in a safe environment.
5. Look for large sequential scans, expensive sorts, poor row estimates, heavy joins, or lock waiting.
6. Fix the actual cause: an index, bounded result set, better query shape, shorter transaction, or additional capacity.
7. Measure again after the change.

Do not add an index blindly. Indexes improve some reads but cost storage and slow writes.

## Transactions, locks, and incidents

Long transactions keep connections busy and may hold locks. This causes pool waits, slow requests, and eventually timeouts.

- Keep transactions short.
- Do not make network calls while a transaction is open.
- Lock related rows in a consistent order.
- Record deadlocks and lock-timeout errors separately from ordinary validation errors.
- Retry only the whole transaction, and only if the operation is safe to retry.

## Backups and recovery

A backup is useful only if it can be restored.

| Term | Practical question |
| --- | --- |
| RPO (recovery point objective) | How much recent data can we afford to lose? |
| RTO (recovery time objective) | How long can recovery take? |

- Define backup frequency, retention, encryption, and access control.
- Practice restoring into a separate environment.
- Know whether you need whole-database, single-table, or point-in-time recovery.
- Treat migrations and backups differently: migrations change schema; backups recover data.

## Security and access

Use separate database roles:

```text
application role → only the reads/writes the application needs
migration role   → schema changes
admin role       → administration
```

Never run a normal web application as a superuser. Store connection credentials in managed secrets, rotate them, and use encrypted connections where required.

## A practical alert starting point

Alert when an issue persists long enough to affect users, not for one brief spike:

- `p95` query or HTTP latency exceeds the service objective.
- Database timeout/error rate increases.
- Pool wait rate rises while `InUse` stays near the configured maximum.
- Database CPU/IO remains saturated.
- Deadlocks, lock waits, or replication lag cross a product-relevant threshold.
- Backup jobs fail or restore verification is overdue.

Set thresholds from a baseline of normal traffic and the user-facing service objective. A dashboard helps investigation; alerts should call attention to action-worthy conditions.

## Incident checklist

1. Is the database reachable? Is the error rate increasing?
2. Is the application waiting for pool connections?
3. Is query latency, CPU, disk IO, or lock waiting high?
4. Did a deployment, migration, traffic spike, or scheduled job coincide with the problem?
5. Mitigate safely: reduce load, stop a harmful job, roll back a bad query/deployment, or scale according to the runbook.
6. After recovery, document the root cause and add a test, metric, limit, index, or alert that prevents recurrence.

## Interview answer

“For production databases, I measure operation latency at percentiles, errors, pool waits, connection usage, resource saturation, and lock contention. I use context deadlines, tune the connection pool from a real connection budget, diagnose slow queries with `EXPLAIN ANALYZE`, use least-privilege access, and verify backups by restoring them.”
