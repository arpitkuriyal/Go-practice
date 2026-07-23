# 07. Concurrent Maps: Share a Map Safely

## The problem

A normal Go map is safe when one goroutine uses it. It is not safe when goroutines read and write it at the same time.

```go
counts := make(map[string]int)
go func() { counts["go"]++ }()
go func() { counts["go"]++ }()
```

This can cause a race or a `concurrent map writes` crash.

## Start with a mutex

For most programs, put a map and a mutex together in one struct:

```go
type Store struct {
	mu sync.Mutex
	m  map[string]int
}

func (s *Store) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}
```

Only one goroutine can hold a `Mutex` at a time. That makes the map update safe.

## Reading safely too

If a writer can run at the same time, reads also need protection:

```go
func (s *Store) Get(key string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.m[key]
	return value, ok
}
```

Use `sync.RWMutex` later when you have many independent reads and few writes:

```go
mu.RLock()
value := m[key]
mu.RUnlock()
```

`RWMutex` allows several readers together, but only one writer. Start with `Mutex` unless measurement shows read-heavy contention.

## `WaitGroup` is not a map lock

```go
wg.Wait() // waits for goroutines to finish
```

It does not stop goroutines from changing a map at the same time. Use it together with a mutex when you need both waiting and safety.

## What about `sync.Map`?

`sync.Map` is a special concurrent map. It can be useful for read-mostly data or independent keys written by many goroutines. For normal application state, `map` plus a mutex is easier to read, type-safe, and better for rules involving more than one key.

## Rules to remember

- Lock the whole rule, not only one line. A “read balance, check, update balance” sequence needs one lock around all three steps.
- Do not copy a struct containing a mutex after using it.
- Keep locked sections short; do not do slow network/database work while holding a lock.
- Run `go test -race ./...` to find races the program might not visibly crash on.

## Interview answer

“Built-in maps are not safe for concurrent access when writes happen. I normally keep a map behind a mutex, lock both reads and writes when they can overlap, and choose `RWMutex` or `sync.Map` only for a measured need.”

## Run

```bash
go run ./concepts/07-concurrent-maps
```
