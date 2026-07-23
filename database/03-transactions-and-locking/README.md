# 03. Transactions and Locking: Keep Related Changes Correct

## The beginner problem

Imagine transferring ₹5 from account A to account B:

```text
1. subtract ₹5 from A
2. add ₹5 to B
```

If the program crashes after step 1, money disappears. A **transaction** groups related statements so they all succeed together or all get undone.

```text
before: A = 100, B = 50
success: A = 95,  B = 55
failure: A = 100, B = 50
```

## The three commands

```sql
BEGIN;    -- start a transaction
COMMIT;   -- make every change permanent
ROLLBACK; -- undo every uncommitted change
```

See [`transfer.sql`](transfer.sql) for a small PostgreSQL example.

## ACID in plain language

| Letter | Meaning |
| --- | --- |
| Atomicity | All steps happen or none happen. |
| Consistency | Database rules remain true after the transaction. |
| Isolation | Concurrent work does not see unsafe half-finished changes. |
| Durability | A committed change survives a crash. |

You do not need to memorize the words first. Remember the money-transfer example: a transaction protects a business rule across multiple queries.

## Go transaction pattern

```go
tx, err := db.BeginTx(ctx, nil)
if err != nil {
	return err
}
defer tx.Rollback() // harmless after a successful commit

if _, err := tx.ExecContext(ctx, debitQuery, amount, fromID); err != nil {
	return err
}
if _, err := tx.ExecContext(ctx, creditQuery, amount, toID); err != nil {
	return err
}
return tx.Commit()
```

The important rule is: after `BeginTx`, use `tx.ExecContext` and `tx.QueryContext` for every statement that belongs to the transaction. Using `db.ExecContext` by mistake runs outside it.

## Why locking exists

Two transfers can read the same balance at nearly the same time. Without coordination, both might decide there is enough money and overspend the account.

PostgreSQL can lock selected rows:

```sql
SELECT id, balance
FROM accounts
WHERE id IN ($1, $2)
ORDER BY id
FOR UPDATE;
```

`FOR UPDATE` makes another transaction wait before changing those rows. Lock accounts in a consistent order (such as increasing ID) to reduce deadlocks.

## Rules to remember

- Keep transactions short: do not call an external API or wait for user input while holding locks.
- Validate the business rule inside the transaction when concurrent updates matter.
- Deadlocks can still happen; a database may abort one transaction. Retry the whole transaction only when the operation is safe to retry.
- `Read Committed` is a common default isolation level. Stronger isolation is useful only when a real invariant needs it.

## Interview answer

“I use a transaction when several database changes represent one business action. Every statement uses the transaction handle, I roll back on failure, commit only after all steps pass, and lock rows when a concurrent read-then-write invariant must hold.”
