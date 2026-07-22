# Concurrency: Revision

## Quick Revision

- **Goroutines** run concurrently with the main goroutine.
- **The program exits** when the main goroutine exits.
- **Use `sync.WaitGroup`** to wait for goroutines.
- **Don't use `time.Sleep`** for synchronization.
- **Go 1.22+** creates a new loop variable for each iteration.
- **Unbuffered channels** require a sender and receiver at the same time.
- **Buffered channels** can store values up to their capacity.
- **`range` on a channel** stops only after the channel is closed.
- **`x++` is not atomic;** protect shared data with a `sync.Mutex`.
- **Use `go test -race`** to detect data races.

---

## Core tools

| Tool | Use it for |
| --- | --- |
| Goroutine | Concurrent unit of work: `go work()`. |
| `sync.WaitGroup` | Waiting for a known set of goroutines. Call `Add` before starting them. |
| Channel | Communicating ownership or results between goroutines. |
| `select` | Waiting on multiple channel operations, cancellation, or timeouts. |
| `sync.Mutex` | Protecting shared mutable state. |
| `context.Context` | Cancellation, deadlines, and request-scoped metadata. |

## Must-know rules

- The program exits when `main` returns.
- Use `sync.WaitGroup`, not `time.Sleep`, to wait for goroutines.
- An unbuffered channel requires both a sender and a receiver.
- A buffered channel blocks only when it is full (send) or empty (receive).
- Close a channel only from the sending side, and only when no more values will be sent.
- `range` over a channel exits only after the channel is closed.
- Protect shared mutable state with `sync.Mutex` or another synchronization primitive.
- Run concurrent code with `go test -race ./...` to detect data races.

## Modern Go loop note

Go 1.22+ creates a fresh loop variable for each iteration declared with `:=`.

Passing the value explicitly is still a good habit because it's clear and works in every Go version:

```go
for _, job := range jobs {
	go func(job Job) {
		process(job)
	}(job)
}
```

## Interview answer: Channel or Mutex?

- Use **`sync.Mutex`** when multiple goroutines share and modify the same data.
- Use **channels** to communicate values, distribute work, or build pipelines.
- Prefer the simpler design—channels are not automatically better than mutexes.

Run the examples:

```bash
go run ./concepts/03-concurrency
```