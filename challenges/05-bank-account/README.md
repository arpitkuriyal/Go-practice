# Bank Account with Error Handling

An in-memory account that validates deposits and withdrawals. `Account` uses a mutex so that balance updates are safe when called concurrently.

Run the tests:

```bash
go test -race ./challenges/05-bank-account
```
