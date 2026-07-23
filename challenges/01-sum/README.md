# Sum Numbers

## Challenge

Write functions that sum two numbers, three numbers, and any number of integers.

```go
sum2(2, 3)          // 5
sum3(1, 2, 3)       // 6
sum(1, 2, 3, 4, 5) // 15
```

## Core idea: variadic functions

`...int` lets a function receive zero or more integer arguments. Inside the function, `nums` is a `[]int`.

```go
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
```

When a slice already exists, expand it with `...` at the call site:

```go
numbers := []int{1, 2, 3}
total := sum(numbers...)
```

## `...` has two meanings

| Location | Meaning |
| --- | --- |
| Function parameter: `func sum(nums ...int)` | Accept many `int` arguments. |
| Function call: `sum(numbers...)` | Expand a slice into individual arguments. |

## Complexity and edge cases

- Time: `O(n)` for `n` values.
- Extra space: `O(1)` besides the variadic argument slice.
- `sum()` returns `0`, the additive identity.
- Integer overflow is possible when the result exceeds the chosen integer type.

## Interview answer

“A variadic parameter is received as a slice. I initialize an accumulator to zero, range over the values, and use `slice...` when I need to pass an existing slice to the function.”

## Run

```bash
go run ./challenges/01-sum/sum.go
```
