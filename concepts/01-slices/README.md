# Slices: Revision

## Mental model

A slice is a small header pointing at an underlying array:

```go
type sliceHeader struct {
    data *T
    len  int
    cap  int
}
```

Copying a slice copies the header, not its elements. Two slices can therefore change the same backing array.

## Rules to remember

| Topic | Rule |
| --- | --- |
| `len(s)` | Number of accessible elements. |
| `cap(s)` | Elements available before `append` must allocate a new backing array. |
| `append` | Always use its returned slice: `s = append(s, value)`. |
| Function argument | A slice header is passed by value, but its elements remain shared. |
| Copy safely | `clone := append([]T(nil), original...)` or `slices.Clone(original)`. |
| Delete | `s = append(s[:i], s[i+1:]...)`; clear removed reference slots when memory retention matters. |

## Interview traps

- `append` can overwrite values visible through another slice when spare capacity exists.
- Appending inside a function does not update the caller's slice length unless the function returns the new slice.
- `nil` and empty slices both have length zero, but `nil` is `nil`; JSON commonly encodes them as `null` and `[]` respectively.
- Keeping a tiny subslice of a very large slice can keep the whole backing array alive. Copy the needed portion when appropriate.

## Say this in an interview

“Slices are descriptors over arrays. Element changes are visible through aliases. `append` may allocate, so I always retain its return value and avoid relying on backing-array sharing.”

Run the examples:

```bash
go run ./concepts/01-slices
```
