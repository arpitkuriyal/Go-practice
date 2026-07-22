# Maps: Revision

## Quick Revision

- **Nil map:** Read ✓ | Write ✗ (panic)
- **Initialize:** `make(map[K]V)` or `map[K]V{}`
- **Maps share** the same underlying hash table.
- **Passing a map** copies only the map header.
- **Modifying entries** affects the caller.
- **Reassigning the map** inside a function doesn't affect the caller.
- **delete()** is always safe, even if the key doesn't exist.
- **Key existence:** `value, ok := m[key]`
- **Iteration order** is intentionally random.
- **Maps are not safe** for concurrent reads and writes without synchronization.

---

## Essential rules

| Operation | Behaviour |
| --- | --- |
| Read from a nil map | Safe; returns the value type's zero value. |
| Write to a nil map | Panics. Initialize with `make` or a literal first. |
| Missing key | Returns zero value; use `value, ok := m[key]` when zero is meaningful. |
| Delete missing key | Safe no-op: `delete(m, key)`. |
| Iteration | Deliberately unspecified order; sort keys if order matters. |
| Equality | Maps cannot be compared except with `nil`. |
| Concurrency | A normal map is not safe for concurrent reads and writes. |

## Useful patterns

```go
counts[word]++                         // zero value makes counters easy
groups[key] = append(groups[key], id)  // append works with a nil slice

value, ok := settings["timeout"]
if !ok { /* choose a default */ }
```

## Concurrency choice

- Use `map` plus `sync.RWMutex` for most shared state: clear typing and invariants.
- Use `sync.Map` only for specialized cases, such as read-mostly keys or independent per-key entries.
- Never rely on a map surviving concurrent access without synchronization; use `go test -race` to find races.

## Interview traps

- A map passed to a function can have its entries changed, but assigning a new map to the parameter does not replace the caller's map variable.
- Map values are not addressable: update a struct value by reading it, modifying a copy, then assigning it back.
- A nil map works well for read-only optional state, but not for lazy writes without initialization.

Run the examples:

```bash
go run ./concepts/02-maps
```