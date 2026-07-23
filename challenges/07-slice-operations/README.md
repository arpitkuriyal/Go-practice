# Generic Slice Operations

## Challenge

Implement generic helpers to insert, remove, filter, and deduplicate slices without changing the caller’s backing array.

## Why return a new slice?

Slices can share an underlying array. Using `append(items[:i], ...)` directly may overwrite values visible through another slice. These helpers allocate a result first, so the input remains unchanged.

```go
result := make([]T, 0, len(items)+1)
result = append(result, items[:index]...)
result = append(result, value)
result = append(result, items[index:]...)
```

## Generic constraints

| Function | Constraint | Why |
| --- | --- | --- |
| `Insert`, `Remove`, `Filter` | `T any` | They do not compare elements. |
| `Unique` | `T comparable` | Map keys must be comparable. |

## Complexity

| Function | Time | Extra space |
| --- | --- | --- |
| Insert / Remove | `O(n)` | `O(n)` |
| Filter | `O(n)` | `O(n)` |
| Unique | `O(n)` expected | `O(n)` |

`Unique` uses a `map[T]struct{}` as a set and preserves the first occurrence order.

## Edge cases

- Insert accepts an index from `0` through `len(items)`.
- Remove requires an existing index from `0` through `len(items)-1`.
- Invalid indexes return `ErrIndexOutOfRange`.
- A nil input is safe; each helper returns a valid empty/new result according to its operation.
- Slices, maps, and functions are not `comparable`, so they cannot be passed to `Unique`.

## Interview answer

“I use generics when the algorithm is independent of element type. I allocate a result to avoid backing-array aliasing, and constrain deduplication to `comparable` because it uses a map.”

## Test

```bash
go test ./challenges/07-slice-operations
```
