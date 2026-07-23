# 02. Slices: Flexible Lists in Go

 

A slice is a list whose length can grow or shrink.

```go
numbers := []int{10, 20, 30}
numbers = append(numbers, 40)
fmt.Println(numbers) // [10 20 30 40]
```

Use a slice when you do not know the exact number of items ahead of time.

## Array versus slice

```go
array := [3]int{1, 2, 3} // always exactly 3 items
slice := []int{1, 2, 3}  // can grow with append
```

| Array | Slice |
| --- | --- |
| Fixed size. | Flexible length. |
| Size is part of its type: `[3]int`. | Size is not part of its type: `[]int`. |
| Less common for application lists. | Used constantly in Go programs. |

## `len` and `cap`

```go
items := make([]int, 3, 5)
fmt.Println(len(items)) // 3: items you can use now
fmt.Println(cap(items)) // 5: room before Go needs a new array
```

A slice is a small description of an underlying array: where it starts, how many values are visible (`len`), and how much room it has (`cap`). You normally do not need to think about this until sharing and `append` matter.

## The most important `append` rule

Always keep the result:

```go
items = append(items, 40)
```

`append` may reuse the old backing array or create a new larger one. Either way, it returns the slice you should use next.

## Slices can share data

```go
original := []int{1, 2, 3}
part := original[:2]
part[0] = 99

fmt.Println(original) // [99 2 3]
```

`part` and `original` point at the same underlying values. This is fast, but it can surprise you.

To make an independent copy:

```go
copyOfItems := append([]int(nil), original...)
```

## Passing a slice to a function

The slice description is copied, but its visible elements are usually shared.

```go
func changeFirst(items []int) {
	items[0] = 100 // caller sees this change
}
```

Appending inside a function does not update the caller’s slice length unless you return the new slice:

```go
func add(items []int, value int) []int {
	return append(items, value)
}
```

## Common traps

- `nil` slice and empty slice both have length zero, but `nil` is `nil`. JSON often encodes them as `null` and `[]`.
- Removing an item with `append(items[:i], items[i+1:]...)` can change the backing array. Copy first when callers must not observe changes.
- Keeping a tiny part of a huge slice can keep the large array in memory. Copy the small part when necessary.

## Interview answer

“A slice is a flexible view over an underlying array. `len` is the visible size and `cap` is available room. I always keep the result of `append` and remember that slices can share their underlying values.”

## Run

```bash
go run ./concepts/02-slices
```
