# Bank Account: Errors and Safe Shared State

## Challenge

Implement an in-memory account that validates deposits and withdrawals, reports useful errors, and remains correct when many goroutines use it at once.

## Concepts practised

| Concept | Use in this challenge |
| --- | --- |
| Sentinel errors | Callers check invalid amounts and insufficient funds with `errors.Is`. |
| Error wrapping | A withdrawal adds the attempted amount while preserving its cause. |
| `sync.Mutex` | Protects balance reads and updates. |
| Invariant | Balance must never become negative. |

## Core pattern

The check and update belong in the same critical section:

```go
a.mu.Lock()
defer a.mu.Unlock()

if amount > a.balance {
	return fmt.Errorf("withdraw %.2f: %w", amount, ErrInsufficientFunds)
}
a.balance -= amount
```

Locking only the subtraction is not enough: another goroutine could change the balance after the check.

## Rules and edge cases

- Reject negative or zero deposits and withdrawals.
- Reject a negative initial balance.
- Use `errors.Is(err, ErrInsufficientFunds)` rather than comparing a wrapped error with `==`.
- Keep all balance access synchronized, including reads.
- `float64` is convenient for practice but not ideal for real money; production systems often store integer minor units (for example, cents) or use a decimal type.

## Complexity

Each operation is `O(1)`. A `Mutex` makes individual operations safe, but it serializes concurrent balance access.

## Interview answer

“I protect the account invariant with one mutex around the whole read-check-update sequence. Expected failures use sentinel errors, and wrapped errors retain context without preventing `errors.Is` checks.”

## Test

```bash
go test -race ./challenges/05-bank-account
```
