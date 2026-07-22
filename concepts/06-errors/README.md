# Errors: Revision

## Mental model

An error is an interface value. Functions return `nil` on success and a non-nil error on failure. Add useful operation context as an error moves up the call stack, while preserving a cause callers need to handle.

```go
var ErrInvalidEmail = errors.New("invalid email")

return fmt.Errorf("register user: %w", ErrInvalidEmail)
```

`%w` preserves `ErrInvalidEmail` in the error chain, so the extra message does not prevent correct handling.

## Error tools

| Tool | Use it for | Example |
| --- | --- | --- |
| `errors.New` | A stable sentinel condition. | `var ErrInvalidEmail = errors.New("invalid email")` |
| `fmt.Errorf` with `%w` | Adding operation context and preserving a cause. | `fmt.Errorf("register user: %w", err)` |
| `errors.Is` | Checking a known error anywhere in a wrapped chain. | `errors.Is(err, ErrInvalidEmail)` |
| `errors.As` | Extracting a typed error and its fields. | `errors.As(err, &validationErr)` |

## Sentinel versus typed errors

Use a **sentinel error** when the caller only needs the category of failure:

```go
if errors.Is(err, ErrInvalidEmail) {
	// Tell the client to correct the email address.
}
```

Use a **typed error** when the caller needs structured details. `ValidateAge` returns `*ValidationError`, so callers can inspect the invalid field and reason:

```go
var validationErr *ValidationError
if errors.As(err, &validationErr) {
	fmt.Println(validationErr.Field, validationErr.Reason)
}
```

## Rules to remember

- Do not use `err == ErrInvalidEmail` after wrapping; use `errors.Is`.
- Wrap at a meaningful operation boundary: `"load user: %w"` is useful; `"failed: %w"` is usually not.
- Use `%w` only when the caller should inspect the original cause. Use `%v` when intentionally hiding implementation details.
- Return errors upward; normally log them once at the application boundary.
- Do not use `panic` for ordinary input validation, missing records, or network failures.
- Keep public error messages free from passwords, tokens, and internal infrastructure details.

## Interview answer: `errors.Is` or `errors.As`?

Use `errors.Is` to ask whether an error represents a known condition. Use `errors.As` when you need a particular error type and the data it contains.

## Run the tests

```bash
go test ./concepts/06-errors
```
