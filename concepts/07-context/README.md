# Context: Revision

## Mental model

`context.Context` carries cancellation, deadlines, and request-scoped values across API boundaries. It does not replace normal function parameters or become a bag for optional values.

```go
ctx, cancel := context.WithTimeout(parent, time.Second)
defer cancel()

err := Wait(ctx, ready)
```

Always call the returned `cancel`, even if the deadline may not fire. It releases resources associated with the derived context.

## Rules to remember

| Topic | Rule |
| --- | --- |
| Parameter position | Put `context.Context` first: `func Load(ctx context.Context, id string)`. |
| Parent context | Propagate the caller's context; do not replace it with `context.Background()` during request work. |
| Timeouts | Derive them with `context.WithTimeout` and always call `cancel`. |
| Cancellation | Check `ctx.Done()` in work that can block or run a long time. |
| Error | Return `ctx.Err()` after cancellation: `context.Canceled` or `context.DeadlineExceeded`. |
| Values | Use only for request-scoped cross-cutting data, such as a trace ID, with a private typed key. |

## Cancellation-aware channel work

`Stream` sends values from a goroutine. A send can block forever if a consumer stops early, so the send must also observe cancellation:

```go
select {
case out <- value:
case <-ctx.Done():
	return
}
```

The producer owns `out`, so it closes it. Consumers should not close a channel they only receive from. `Wait` uses the same pattern for a blocking receive: it returns normally when work is ready or returns the context error if cancellation wins.

## Avoiding goroutine leaks

Every goroutine needs a clear exit condition: completed work, a closed input channel, a delivered result, or context cancellation. Make cancellation observable at every blocking send and receive. If a consumer can stop early, it should cancel the shared context so upstream producers can exit.

## Interview answers

### Why is `context.Context` the first argument?

It makes request lifetime control visible in the signature and follows the standard-library convention. The caller can then cancel the full call chain consistently.

### Who closes a channel?

The sending side that knows no more values will be sent. Closing from a receiver risks a send-on-closed-channel panic.

## Run the tests

```bash
go test -race ./concepts/07-context
```
