# Go Practice and Interview Revision

This repository is a hands-on Go revision workspace. It contains short programs, notes, and challenges for strengthening Go fundamentals before technical interviews.

## Topics covered

| Area | What you will find |
| --- | --- |
| Core Go | Strings, bytes and runes, slices, maps, structs, methods, interfaces, and generics |
| Errors | Detailed examples for sentinel errors, wrapping, `errors.Is`, `errors.As`, and typed errors |
| Context | Detailed examples for cancellation, deadlines, context-aware channels, and goroutine-leak prevention |
| Concurrency | Goroutines, channels, `select`, mutexes, `WaitGroup`, concurrent maps, worker pools, semaphores, rate limiting, caches, `sync.Once`, and `errgroup` |
| Testing and tooling | Table-driven tests, subtests, fakes, benchmarks, fuzzing, race detection, `go vet`, modules, and formatting |
| HTTP and APIs | A beginner-to-production revision path covering request/response flow, methods, headers, statuses, routing, middleware, JSON APIs, safe input handling, timeouts, testing, and graceful shutdown |
| Databases and SQL | `database/sql`, parameterized queries, connection pools, migrations, transactions, locking, indexes, PostgreSQL features, and repository testing |
| Challenges | Nine progressively harder problems covering Go basics, data handling, interfaces, concurrency, strings, error handling, and HTTP middleware |

## Repository layout

```text
.
├── challenges/
│   ├── 01-sum/                 # Basic arithmetic exercise
│   ├── 02-reverse-string/      # Strings and rune exercise
│   ├── 03-employee-management/ # Structs and slice manipulation
│   ├── 04-word-frequency/      # Maps and text processing
│   ├── 05-bank-account/        # Errors and concurrent state
│   ├── 06-shape-calculator/    # Interfaces and polymorphism
│   ├── 07-slice-operations/    # Generic slice helpers
│   ├── 08-string-pattern-matching/ # Substring and anagram matching
│   └── 09-http-auth-middleware/    # HTTP middleware and context
├── concepts/
│   ├── 01-slices/
│   ├── 02-maps/
│   ├── 03-concurrency/
│   ├── 04-concurrent-maps/
│   ├── 05-interfaces/
│   ├── 06-errors/
│   ├── 07-context/
│   ├── 08-testing-and-tooling/
│   └── 09-concurrency-patterns/
├── database/
│   ├── 01-sql-basics/          # SQL, parameterized queries, and pooling
│   ├── 02-migrations/          # Versioned schema changes and constraints
│   ├── 03-transactions-and-locking/
│   ├── 04-query-performance/
│   ├── 05-postgresql-features/
│   └── 06-repository-testing/
└── http/
    ├── 01-basics/
    ├── 02-methods/
    ├── 03-requests-and-responses/
    ├── 04-server-and-routing/
    ├── 05-middleware/
    ├── 06-json-encoding/
    └── 07-production-checklist/ # Production HTTP patterns and tested API
```

## Running an example

Run an individual Go file from the repository root:

```bash
go run ./challenges/04-word-frequency/word-frequency.go
```

For a directory containing one runnable package, you can also use:

```bash
go run ./concepts/05-interfaces
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

## Contributing

Contributions that improve the Go notes, examples, tests, and interview revision material are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request. This repository is available under the [MIT License](LICENSE).

## Featured practice challenges

Use this list as a question bank. First read only the **challenge** column and solve it yourself. If you get stuck, open the linked **solution** in this repository and compare your approach after you finish.

| # | Challenge | What to build | Solution |
| --- | --- | --- | --- |
| 01 | Sum numbers | Write fixed-arity and variadic functions that return the sum of integer inputs. | [Open solution](challenges/01-sum/sum.go) |
| 02 | Reverse a string | Reverse a string without modifying it in place; then consider how Unicode changes the solution. | [Open solution](challenges/02-reverse-string/reverse.go) |
| 03 | Employee management | Create an employee manager that adds, removes, finds employees by ID, and calculates average salary. | [Open solution](challenges/03-employee-management/data-management.go) |
| 04 | Word frequency | Normalize a sentence and return a map of each word to its number of occurrences. | [Open solution](challenges/04-word-frequency/word-frequency.go) |
| 05 | Bank account | Implement validated deposits, withdrawals, balance reads, useful errors, and safe concurrent updates. | [Open solution](challenges/05-bank-account/account.go) |
| 06 | Shape calculator | Define a `Shape` interface for circles, rectangles, and triangles; calculate total area. | [Open solution](challenges/06-shape-calculator/shapes.go) |
| 07 | Slice operations | Implement generic insert, remove, filter, and deduplication helpers without mutating the input. | [Open solution](challenges/07-slice-operations/slices.go) |
| 08 | String pattern matching | Find overlapping substring matches, then implement Unicode-safe anagram matching with a sliding window. | [Open solution](challenges/08-string-pattern-matching/patterns.go) |
| 09 | HTTP authentication middleware | Validate a Bearer token, return JSON errors, and attach the authenticated user to request context. | [Open solution](challenges/09-http-auth-middleware/middleware.go) |

The new challenges include unit tests; use them only after making your own attempt. Run all tests with `go test ./...`. For challenges 05 and 09, also run `go test -race ./...`.

## Useful interview prompts

- Why does `append` sometimes change the original slice and sometimes not?
- When would you use a channel instead of a mutex?
- What happens when a goroutine sends on an unbuffered channel?
- How does interface satisfaction work in Go?
- How would you avoid goroutine leaks in an HTTP request?
- How do you distinguish an expected error from an exceptional failure?
