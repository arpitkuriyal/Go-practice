# Concurrency: Revision

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

- The program exits when `main` returns; do not use `time.Sleep` as synchronization.
- An unbuffered send/receive rendezvous; a buffered channel blocks only when it is full (send) or empty (receive).
- Close a channel only from the sending side, and only when no more values will be sent.
- Ranging over a channel ends only after the channel is closed.
- A goroutine that cannot make progress is a goroutine leak. Every blocking operation should have an owner or cancellation path.
- Run concurrent code with `go test -race ./...`; the race detector finds unsynchronized memory access, not every logical bug.

## Modern Go loop note

For loops declared with `:=` now create a fresh loop variable per iteration in Go 1.22+. Passing the value explicitly remains clear and works with every supported style:

```go
for _, job := range jobs {
    go func(job Job) { process(job) }(job)
}
```

## Interview answer: channel or mutex?

Use a mutex when goroutines need shared state with simple invariants. Use channels when transferring work, applying backpressure, or coordinating a pipeline. Prefer the simpler design—channels are not automatically safer.

Run the examples:

```bash
go run ./concepts/03-concurrency
```
