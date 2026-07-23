# 06. Repository Testing: Prove Your SQL Works

## Start simple: what is a repository?

A repository is the small Go layer that runs SQL and converts database rows into Go values.

```text
HTTP handler / service → repository → PostgreSQL
```

For example, a repository method might execute `SELECT ... FROM users WHERE id = $1`, scan the result into a `User`, and return it.

## Why test it with a real database?

A fake repository can prove that service logic calls a method correctly. It cannot prove that SQL is valid, placeholders match the driver, columns scan correctly, or constraints behave as expected.

Use both kinds of tests:

| Test type | What it proves |
| --- | --- |
| Unit test with a fake | Business/service logic is correct and fast. |
| Repository integration test | Real SQL, migrations, scans, constraints, and transactions work. |
| End-to-end test | HTTP through database boundaries works together. |

## A simple integration-test flow

1. Start a dedicated disposable PostgreSQL database (often Docker or Testcontainers).
2. Apply the same migrations used by the application.
3. Insert only the data needed for this test.
4. Call the repository with a timeout context.
5. Assert the returned value or error.
6. Clean up by rollback, truncation, or deleting the disposable database.

Never point tests at a shared or production database.

## Important cases to test

- A missing row maps from `sql.ErrNoRows` to the expected application result.
- A duplicate value violates the intended `UNIQUE` constraint.
- A bad foreign key is rejected.
- A transaction rolls back when a later statement fails.
- Nullable columns and time zones scan correctly.
- Pagination returns a stable, bounded order.
- Concurrent updates preserve the business rule.

## Test data rule

Each test should create its own data with clear values. Do not rely on rows left behind by another test. Independent data makes tests deterministic and safe to run in any order.

## Interview answer

“I unit-test services with fakes for speed, but I integration-test repositories against a disposable real database because only a real database proves SQL syntax, constraints, scanning, and transactions.”
