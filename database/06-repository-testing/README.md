# Repository Testing: Revision

## Test layers

| Layer | What it verifies |
| --- | --- |
| Unit test | Service logic using a fake repository. Fast and deterministic. |
| Repository integration test | Real SQL, driver behaviour, scans, constraints, and transactions. |
| End-to-end test | HTTP handler through database and external boundaries. Use sparingly. |

## Integration-test pattern

1. Start a dedicated disposable database, commonly with Docker or Testcontainers.
2. Apply the same migrations used by the application.
3. Create test data only for the case being tested.
4. Run the repository call with a timeout context.
5. Clean up with a transaction rollback, truncation, or a newly created database.

## What to verify

- `sql.ErrNoRows` maps to the expected domain outcome.
- Unique and foreign-key violations are handled correctly.
- Transactions roll back when a later statement fails.
- Query results scan the right nullable values and time zones.
- Pagination, ordering, and concurrent updates preserve the intended invariant.

Avoid testing SQL by asserting only query strings in unit tests. A fake can prove service behaviour, but a real database test is what proves the query works.
