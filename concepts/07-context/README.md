# 07. Context: Stop Work That No Longer Matters

A context lets one part of a program tell another part: “stop this work,” “there is a deadline,” or “this request has a trace ID.”

The most common use is cancelling database, HTTP, or goroutine work when a user leaves a page or a timeout is reached.

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()

err := doWork(ctx)
```

`cancel` should always be called. It releases resources even when the timeout never happens.

## Put context first

```go
func LoadUser(ctx context.Context, id string) error
```

Context is the first parameter by Go convention. Accept the caller’s context and pass it to the next operation:

```go
row := db.QueryRowContext(ctx, query, id)
```

Do not replace request work with `context.Background()`: that throws away the caller’s cancellation and deadline.

## Notice cancellation

`ctx.Done()` is a channel that becomes ready when the context is cancelled or expires:

```go
select {
case <-ctx.Done():
	return ctx.Err()
case result := <-results:
	return result
}
```

`ctx.Err()` is normally `context.Canceled` or `context.DeadlineExceeded`.

## Why goroutines need context

A goroutine can get stuck forever when it tries to send on a channel that nobody reads. The `Stream` example avoids that by watching both the send and cancellation:

```go
select {
case out <- value:
	// consumer received the value
case <-ctx.Done():
	return // caller stopped caring
}
```

The sender closes the channel because it knows when no more values will be sent. A receiver should not close a channel it only receives from.

## Context values

Use values only for small request-scoped information that crosses layers, such as a request ID or authenticated user. Use a private typed key. Do not use context as a bag for optional function arguments, database handles, or configuration.

## Rules to remember

- Pass context to blocking work: database queries, outbound HTTP calls, and goroutines.
- Always call the cancel function returned by `WithCancel`, `WithTimeout`, or `WithDeadline`.
- Every goroutine needs an exit path: completed work, closed input, sent result, or cancellation.
- A timeout is not a replacement for making slow code efficient; it is a safety boundary.

## Interview answer

“Context carries cancellation and deadlines through a call chain. I accept it as the first parameter, propagate it to blocking work, call cancel for derived contexts, and make goroutines observe `ctx.Done()` so they do not leak.”

## Test

```bash
go test -race ./concepts/07-context
```
