# Database and SQL: Revision

## `database/sql` rules

- `sql.DB` is a concurrency-safe connection pool, not one database connection. Create it once and share it.
- Use `QueryRowContext`, `QueryContext`, and `ExecContext` so request cancellation reaches the database driver.
- Always use query placeholders and arguments. Never build SQL by concatenating user input.
- Check `sql.ErrNoRows` with `errors.Is` after a query for one row.
- Close `Rows` with `defer rows.Close()` and check `rows.Err()` after iteration.

## Prepared statements and injection

Arguments passed separately from the SQL text are parameterized by the driver. This prevents a value from changing query syntax:

```go
db.QueryRowContext(ctx,
    "SELECT id, name FROM users WHERE email = $1", email)
```

Use a prepared statement only when it improves a known workload; most drivers and databases already optimize repeated parameterized queries. Placeholder syntax differs: PostgreSQL uses `$1`, while some drivers use `?`.

## Transactions

Use `BeginTx`, execute every related statement through `tx`, then `Commit`. Defer `Rollback` immediately after a successful begin; it becomes a safe no-op after a commit. Keep transactions short and use appropriate isolation/locking for invariants such as account balances.

## Pool, migrations, and indexes

- Set pool limits based on the database's connection capacity and real metrics, not guesses.
- Run migrations in versioned files through a migration tool; do not hide schema changes in application startup.
- Index columns used in `WHERE`, joins, and common ordering only after checking the query plan. Indexes speed reads but cost storage and write time.

`repository.go` is a PostgreSQL-style example. It intentionally does not include a driver or credentials.
