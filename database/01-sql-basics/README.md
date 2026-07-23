# 01. SQL Basics: Store and Read Data

Start here if databases and SQL are new.

## What is a database?

A database stores data after your Go program stops. A relational database stores data in **tables**.

```text
users table
+----+-------+-------------------+
| id | name  | email             |
+----+-------+-------------------+
| 1  | Arpit | arpit@example.com |
| 2  | Neha  | neha@example.com  |
+----+-------+-------------------+
```

- A **table** is like a spreadsheet with named columns.
- A **row** is one record, such as one user.
- A **column** describes one property, such as `email`.
- SQL is the language used to ask the database to create, read, update, and delete rows.

## The four basic operations: CRUD

| Letter | SQL command | Plain meaning |
| --- | --- | --- |
| Create | `INSERT` | Add a row. |
| Read | `SELECT` | Get rows. |
| Update | `UPDATE` | Change existing rows. |
| Delete | `DELETE` | Remove rows. |

Open [`queries.sql`](queries.sql). It contains a complete small example. The key commands are:

```sql
-- Read selected columns from all users.
SELECT id, name, email FROM users;

-- Read only the matching user.
SELECT id, name, email
FROM users
WHERE email = 'arpit@example.com';
```

## Read SQL in order

For a simple query, read it as:

```sql
SELECT id, name       -- which columns?
FROM users            -- which table?
WHERE id = 1;         -- which rows?
```

Always use a `WHERE` clause with `UPDATE` and `DELETE` unless you truly mean every row:

```sql
UPDATE users SET name = 'Arpit' WHERE id = 1;
DELETE FROM users WHERE id = 1;
```

Without `WHERE`, the command affects all rows in the table.

## Table rules: constraints

Constraints protect data even if an application has a bug.

```sql
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL
);
```

| Rule | Meaning |
| --- | --- |
| `PRIMARY KEY` | A unique ID for every row. |
| `NOT NULL` | This column is required. |
| `UNIQUE` | Two rows cannot have the same value, such as an email. |

Validate user-friendly input in Go too, but keep important integrity rules in the database.

## Connect SQL to Go

Go uses the standard `database/sql` package. [`main.go`](main.go) is a repository example; it does not start a database by itself and intentionally has no driver or credentials.

The safest Go pattern is a parameterized query:

```go
row := db.QueryRowContext(ctx,
	"SELECT id, name, email FROM users WHERE email = $1", email)
```

`$1` is a placeholder. The driver sends `email` separately, preventing it from changing the SQL syntax. Never build SQL by joining user input into a query string.

## First Go rules

- `*sql.DB` is a reusable, concurrency-safe connection pool—not one connection per request.
- Create it once when the application starts and share it.
- Pass `r.Context()` from HTTP handlers into `QueryContext`, `QueryRowContext`, and `ExecContext`.
- A query returning no row produces `sql.ErrNoRows`; check it with `errors.Is`.

## Next step

Once CRUD makes sense, learn how schema changes are shared safely with migrations → [02 Migrations](../02-migrations/README.md).

## Interview answer

“A relational database stores rows in tables. I use SQL CRUD operations with `WHERE` clauses, enforce important rules using constraints, and call parameterized context-aware queries through Go’s `database/sql` package.”
