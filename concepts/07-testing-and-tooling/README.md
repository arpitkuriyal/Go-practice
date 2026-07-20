# Testing and Tooling: Revision

## Test styles

| Tool | Best use |
| --- | --- |
| Table-driven test | The same behaviour over many inputs and edge cases. |
| Subtest: `t.Run` | Names each case and allows individual reruns. |
| Fake | A small working test implementation of a dependency, such as an in-memory store. |
| Mock | Verifies interactions; use only when interactions—not output—are the contract. |
| Benchmark | Measures time and allocations; compare before/after changes. |
| Fuzz test | Generates unexpected inputs and keeps any crashing input as a regression case. |

## Commands to know

```bash
go test ./...                         # all unit tests
go test -race ./...                   # detect data races
go test -run TestSum ./path            # one test or matching tests
go test -bench . -benchmem ./path      # benchmarks and allocations
go test -fuzz FuzzNormalizeEmail ./path # fuzz a target
go vet ./...                          # suspicious constructs
gofmt -w .                            # format Go files
go mod tidy                           # add needed and remove unused module requirements
```

## Testing guidance

- Test exported behaviour and important edge cases, not implementation details.
- Keep tests deterministic: do not use real clocks, network calls, or random data without control.
- Depend on a small interface, then use a fake in unit tests. Reserve integration tests for real databases and services.
- Run `gofmt` and `go vet` before committing. Run `go test -race` for shared state and goroutines.

Run this module:

```bash
go test -bench . ./concepts/07-testing-and-tooling
```
