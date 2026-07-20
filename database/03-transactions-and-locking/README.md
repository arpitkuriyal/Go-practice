# Transactions and Locking: Revision

## ACID

| Property | Meaning |
| --- | --- |
| Atomicity | Every statement succeeds, or all changes roll back. |
| Consistency | A transaction preserves defined constraints and invariants. |
| Isolation | Concurrent transactions do not expose unsafe intermediate state. |
| Durability | A committed transaction survives failures. |

## Go pattern

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil { return err }
defer tx.Rollback() // safe after a successful Commit

// Use tx.QueryContext / tx.ExecContext for every related statement.
if _, err := tx.ExecContext(ctx, query, args...); err != nil { return err }
return tx.Commit()
```

## Locking and deadlocks

- Use a transaction when multiple writes must succeed or fail together.
- Lock rows with `SELECT ... FOR UPDATE` when a read-then-write invariant must be protected, for example an account balance.
- Keep transactions short: do not wait on network calls or user input while holding locks.
- Access resources in a consistent order to reduce deadlocks.
- Deadlocks can still happen; PostgreSQL aborts one transaction. Retry only the whole transaction when the operation is safely retryable.

## Isolation interview answer

`Read Committed` is a common default: each query sees committed data. Stronger isolation prevents more anomalies but can increase retries or contention. Choose isolation based on the business invariant, not by always selecting the strongest level.

See the `Transfer` example in [`01-sql-basics`](../01-sql-basics/repository.go).
