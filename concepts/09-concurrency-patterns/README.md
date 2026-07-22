# Concurrency Patterns: Last-Minute Revision

## Worker pool, fan-out, and fan-in

- A **worker pool** has a bounded number of workers reading jobs from one channel. Use it to limit parallelism.
- **Fan-out** sends independent work to several workers. **Fan-in** merges their results into one consumer channel.
- Close the jobs channel after the producer is done; wait for workers before closing a shared results channel.
- Pass `context.Context` through the whole pipeline so a caller can cancel every stage.

## Semaphore and rate limiter

- A buffered channel is a simple semaphore: acquire by sending a token, release by receiving it.
- Always make acquire context-aware and always release after a successful acquire.
- A token-bucket rate limiter allows a short burst, then refills at a fixed rate. Return `429 Too Many Requests` when a request is rejected.

## Cache choices

| Cache | Eviction rule | Good for |
| --- | --- | --- |
| TTL | Remove entries after time expires. | Data with a freshness window. |
| LRU | Remove the least recently used entry at capacity. | Bounded memory with locality of access. |

`patterns.go` includes safe TTL and LRU examples. Production caches also need observability, size limits, and a clear invalidation strategy.

## `sync.Once` and `errgroup`

- Use `sync.Once` for one-time initialization. The function must be safe to run exactly once; do not use it for work that should be retried after failure without an explicit design.
- `errgroup` is from `golang.org/x/sync/errgroup`. It combines a `WaitGroup` with error propagation and is ideal for independent concurrent tasks that should cancel together.

```go
group, ctx := errgroup.WithContext(ctx)
group.Go(func() error { return fetch(ctx, "first") })
group.Go(func() error { return fetch(ctx, "second") })
if err := group.Wait(); err != nil { /* handle first error */ }
```

Install it only when you need it:

```bash
go get golang.org/x/sync/errgroup
```

Run the examples with race detection:

```bash
go test -race ./concepts/09-concurrency-patterns
```
