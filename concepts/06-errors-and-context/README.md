# Errors and Context: Revision

## Errors

```go
var ErrNotFound = errors.New("not found")

return fmt.Errorf("load user %d: %w", id, ErrNotFound)
```

| Tool | Use |
| --- | --- |
| `errors.New` | A stable sentinel error for a condition callers may handle. |
| `%w` | Wrap an error with operation context while preserving the cause. |
| `errors.Is(err, target)` | Check a sentinel error through wrapping. |
| `errors.As(err, &target)` | Extract a typed error, such as `*ValidationError`. |

Do not compare wrapped errors with `==`; use `errors.Is`. Add useful context at the boundary where an operation fails, but avoid repeatedly logging the same error at every layer.

## Context

- Accept `context.Context` as the first parameter of request-scoped functions.
- Derive deadlines with `context.WithTimeout`; always call the returned `cancel`, normally with `defer cancel()`.
- Propagate the same context to SQL queries and outbound HTTP calls.
- Check `<-ctx.Done()` in goroutines and channel sends/receives that could block.
- Do not store a context in a struct or use it to pass optional function parameters.

## Avoiding goroutine leaks

Every goroutine needs a clear completion condition: a closed work channel, a received result, or context cancellation. Every producer should close the channels it owns. Consumers that might stop early should cancel the shared context.

Run the examples and tests:

```bash
go test -race ./concepts/06-errors-and-context
```
