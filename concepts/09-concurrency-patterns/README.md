# 09. Concurrency Patterns: Useful Building Blocks

Learn the basic goroutine, channel, mutex, and context lesson first. These patterns combine those tools for common problems.

## Worker pool: limit parallel work

If you have 1,000 jobs, starting 1,000 goroutines may overwhelm a database or API. A worker pool starts a fixed number of workers and gives them jobs through a channel.

```text
jobs → worker 1 ─┐
     → worker 2 ─┼→ results
     → worker 3 ─┘
```

The included `WorkerPool` limits how many jobs run at once. Use it for batch processing, file work, or calls to a limited external service.

## Semaphore: a small concurrency limit

A semaphore is a limited number of tokens:

```text
3 tokens → at most 3 operations can run
```

The `Semaphore` example uses a buffered channel. `Acquire` gets a token; `Release` returns it. Always release a token after a successful acquire, usually with `defer`.

## Rate limiter: limit how often work starts

Concurrency limit and rate limit are different:

| Tool | Limits |
| --- | --- |
| Semaphore | How many operations run at the same time. |
| Rate limiter | How many operations may start in a time period. |

The example starts with a small burst of tokens and adds a token at each interval. Call `Stop` when finished so its ticker goroutine can exit.

## Cache: remember a result

| Cache | Simple meaning | Good for |
| --- | --- | --- |
| TTL cache | Forget a value after a time limit. | Data that can be slightly old. |
| LRU cache | Remove the least recently used value when full. | A fixed memory limit with repeated lookups. |

The TTL cache removes expired values when `Get` is called. The LRU cache keeps recently used values at the front of a list. Both are protected by a mutex.

Before adding a cache, answer: what is cached, how old can it be, how is it invalidated, and what happens on a miss?

## Rules to remember

- Pass context through worker and semaphore work so callers can cancel.
- Close the jobs channel after the producer finishes sending jobs.
- Wait for workers before closing a shared results channel.
- Keep cache size and lifetime bounded so memory cannot grow forever.
- Test concurrent patterns with `go test -race`.

## Interview answer

“I use a worker pool to bound parallel jobs, a semaphore to limit concurrent resource use, a rate limiter to limit request starts, and a cache when I can define freshness and invalidation. Each pattern needs cancellation and clear cleanup.”

## Test

```bash
go test -race ./concepts/09-concurrency-patterns
```
