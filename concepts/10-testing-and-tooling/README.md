# 10. Testing and Tooling: Check Your Go Code

A Go test is a function in a `_test.go` file:

```go
func TestSum(t *testing.T) {
	if got := Sum(2, 3); got != 5 {
		t.Fatalf("Sum() = %d, want 5", got)
	}
}
```

Run all tests:

```bash
go test ./...
```

Tests give you confidence that code still works after a change.

## Table-driven tests

When the same function has many inputs, keep cases in a table:

```go
tests := []struct {
	name string
	in   []int
	want int
}{
	{"empty", nil, 0},
	{"numbers", []int{1, 2, 3}, 6},
}
```

Then loop through the cases with `t.Run`. This keeps tests short and gives each failure a useful name.

## What should I test?

Test behavior a caller cares about:

- Normal success case.
- Empty input and boundary values.
- Invalid input and expected errors.
- Important business rules.
- A bug you fixed, so it does not return.

Avoid tests that only prove private implementation details. You want freedom to refactor without rewriting every test.

## Fakes and integration tests

A fake is a small working test replacement for a dependency:

```go
type fakeStore struct{ exists bool }
func (f fakeStore) Exists(ctx context.Context, email string) (bool, error) {
	return f.exists, nil
}
```

Use a fake to test service logic quickly. Use an integration test when you need to prove a real database, HTTP service, or driver works correctly.

## Commands to know

```bash
go test ./...                              # all tests
go test -race ./...                        # find data races
go test -run TestSum ./path                # one matching test
go test -bench . -benchmem ./path          # benchmark time and allocations
go test -fuzz FuzzNormalizeEmail ./path    # try generated inputs
go vet ./...                               # find suspicious code
gofmt -w path/to/file.go                   # format a Go file
go mod tidy                                # clean module dependencies
```

## Benchmark and fuzzing in one line

- A benchmark measures speed. Compare before and after a change; do not trust one noisy run.
- A fuzz test generates many strange inputs and saves a failing input as a future regression case.

## Interview answer

“I write small deterministic tests for observable behavior and edge cases. I use table-driven tests for many cases, fakes for fast service tests, integration tests for real boundaries, and the race detector for shared concurrent state.”

## Run this lesson

```bash
go test -bench . ./concepts/10-testing-and-tooling
```
