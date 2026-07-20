# Concurrent Maps: Revision

## The rule

A built-in map is not safe for concurrent access when at least one goroutine writes. It can panic with `concurrent map writes`, and even a read racing with a write is unsafe.

## Default solution: `map` + `sync.RWMutex`

```go
type Store struct {
    mu sync.RWMutex
    m  map[string]int
}

func (s *Store) Get(key string) (int, bool) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    value, ok := s.m[key]
    return value, ok
}

func (s *Store) Set(key string, value int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.m[key] = value
}
```

Use `Mutex` first. Choose `RWMutex` only when profiling or workload knowledge shows many independent reads; it has additional coordination cost.

## `sync.Map`

`sync.Map` is useful for read-mostly data or independent keys written by many goroutines. It is less type-safe and makes multi-step invariants harder to express. For example, “read, check, then update” often still needs another synchronization strategy.

## Interview traps

- Lock reads as well as writes if a writer can run concurrently.
- Do not copy a struct containing a mutex after first use.
- Keep critical sections small, but protect the whole invariant—not just individual lines.
- `LoadOrStore` is useful for per-key initialization; it is not a replacement for a transaction across multiple keys.

Verify with:

```bash
go test -race ./...
go run ./concepts/04-concurrent-maps
```
