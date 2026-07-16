# Go Practice and Interview Revision

This repository is a hands-on Go revision workspace. It contains short programs, notes, and challenges for strengthening Go fundamentals before technical interviews.

## Topics covered

| Area | What you will find |
| --- | --- |
| Strings | String immutability, bytes, runes, and common operations |
| Slices | Slice internals and behaviour |
| Maps | Map internals and map-related examples |
| Concurrency | Concurrency and map-concurrency internals |
| Interfaces | An interval-focused interface exercise |
| HTTP | Basics, methods, headers, status codes, routing, middleware, and encoding/decoding |
| Challenges | Sum, reverse string, employee data management, and word frequency |

## Repository layout

```text
.
├── 001-challenge/              # Sum exercise
├── 002-challenge/              # Reverse-string exercise and notes
├── 003-challenge/              # Employee data-management exercise
├── 004-challenge/              # Word-frequency exercise
├── concurrency-internals/      # Concurrency concepts
├── maps-concurrency-internals/ # Maps used with concurrency
├── map-internals/              # Map concepts
├── slice-internals/            # Slice concepts
├── interfaces-intervals/       # Interfaces exercise
└── http-notes/                 # HTTP notes and examples
```

## Running an example

Run an individual Go file from the repository root:

```bash
go run ./004-challenge/word-frequency.go
```

For a directory containing one runnable package, you can also use:

```bash
go run ./interfaces-intervals
```

Format files before committing changes:

```bash
gofmt -w path/to/file.go
```

## Interview revision checklist

Use this checklist to guide your revision. Aim to explain each item aloud, then write a small example without looking it up.

- [ ] Explain arrays versus slices; include `len`, `cap`, `append`, shared backing arrays, and memory retention.
- [ ] Explain strings, UTF-8, `byte` versus `rune`, and when to use `strings.Builder`.
- [ ] Explain maps, zero values, the comma-ok idiom, and why concurrent map writes are unsafe.
- [ ] Explain structs, methods, pointer receivers, embedding, interfaces, and type assertions/switches.
- [ ] Explain error handling, wrapping with `%w`, `errors.Is`, and `errors.As`.
- [ ] Explain goroutines, channels, `select`, `sync.Mutex`, `sync.WaitGroup`, `context`, race conditions, and deadlocks.
- [ ] Explain `defer`, `panic`, `recover`, escape analysis at a high level, and the Go garbage collector.
- [ ] Build a small HTTP API using `net/http`, JSON, middleware, request validation, and graceful shutdown.
- [ ] Write unit tests, table-driven tests, benchmarks, and run the race detector.

## Suggested workflow

1. Pick one topic and review its notes.
2. Solve or redo one challenge from scratch in 20–30 minutes.
3. Add tests for edge cases before checking your solution.
4. Explain your choices aloud as if answering an interviewer.
5. Run formatting and tests:

   ```bash
   gofmt -w .
   go test ./...
   go test -race ./...
   ```

## Useful interview prompts

- Why does `append` sometimes change the original slice and sometimes not?
- When would you use a channel instead of a mutex?
- What happens when a goroutine sends on an unbuffered channel?
- How does interface satisfaction work in Go?
- How would you avoid goroutine leaks in an HTTP request?
- How do you distinguish an expected error from an exceptional failure?

## Next additions

Add these in roughly this order:

1. **Testing practice** — table-driven tests, subtests, mocks/fakes, benchmarks, fuzz tests, and `go test -race`.
2. **Errors and context** — custom errors, wrapping, cancellation, deadlines, and avoiding goroutine leaks.
3. **Concurrency challenges** — worker pool, rate limiter, fan-out/fan-in, and a thread-safe cache.
4. **A REST API** — CRUD endpoints with JSON, validation, middleware, authentication, graceful shutdown, and tests.
5. **Database work** — `database/sql`, transactions, connection pooling, migrations, and SQL injection prevention.
6. **A CLI tool** — flags, configuration, file I/O, structured logging, and clear error messages.

For every new exercise, include a short problem statement, edge cases, tests, and a note explaining the time/space complexity or design trade-offs.
