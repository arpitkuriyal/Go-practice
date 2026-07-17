# Go Arrays & Slices 🚀

---

# 1. Array vs Slice

| Array | Slice |
|--------|--------|
| Fixed size | Dynamic size |
| Size is part of type | Size isn't part of type |
| Value type | Reference-like type |
| Stored directly | Points to underlying array |

```go
arr := [5]int{1,2,3,4,5}
slice := []int{1,2,3,4,5}
```

---

# 2. Before (`...`) and After (`...`) Operator

The `...` operator has **two meanings** depending on where it is used.

---

## A. Before Type (Variadic Parameter)

Allows a function to accept **0 or more arguments**.

```go
func sum(nums ...int) int {
    total := 0

    for _, n := range nums {
        total += n
    }

    return total
}

sum(1,2,3)
sum(10)
sum()
```

Think of it as

```
1
1,2
1,2,3
1,2,3,4
```

All are valid.

---

## B. After Slice (Expanding Slice)

Suppose a function wants

```go
func sum(nums ...int)
```

and you already have

```go
nums := []int{1,2,3}
```

You **cannot**

```go
sum(nums) ❌
```

because

```
Expected:

1,2,3

Got:

[]int
```

Instead expand it.

```go
sum(nums...)
```

Now Go converts

```
[]int{1,2,3}

↓

1,2,3
```

---

## Another Example

```go
func printNames(names ...string){
    fmt.Println(names)
}

arr := []string{"John","Bob","Alice"}

printNames(arr...)
```

Output

```
[John Bob Alice]
```

---

# Easy Rule

```
Function Definition

...type

means

Accept many values.


Function Call

slice...

means

Expand slice into individual values.
```

---

# 3. Creating Arrays

```go
var a [5]int
```

```
[0 0 0 0 0]
```

---

```go
a := [5]int{1,2,3,4,5}
```

---

```go
a := [...]int{1,2,3}
```

Compiler counts size automatically.

---

# 4. Creating Slices

```go
nums := []int{1,2,3}
```

---

```go
nums := make([]int,5)
```

Length

```
5
```

Capacity

```
5
```

---

```go
nums := make([]int,5,10)
```

Length

```
5
```

Capacity

```
10
```

---

# 5. Length vs Capacity

```go
nums := make([]int,3,5)

fmt.Println(len(nums))
fmt.Println(cap(nums))
```

Output

```
3
5
```

---

# 6. Append

```go
nums := []int{1,2}

nums = append(nums,3)

fmt.Println(nums)
```

Output

```
[1 2 3]
```

Append multiple

```go
nums = append(nums,4,5,6)
```

Append another slice

```go
a := []int{1,2}
b := []int{3,4}

a = append(a,b...)
```

---

# 7. Copy

```go
a := []int{1,2,3}
b := make([]int,len(a))

copy(b,a)
```

Now changing

```go
b[0]=100
```

doesn't affect

```
a
```

---

# 8. Slice Operation

```go
nums := []int{10,20,30,40,50}
```

```
nums[1:4]

↓

20 30 40
```

---

```
nums[:3]

↓

10 20 30
```

---

```
nums[2:]

↓

30 40 50
```

---

# 9. Underlying Array

```go
a := []int{1,2,3}

b := a

b[0]=100

fmt.Println(a)
```

Output

```
[100 2 3]
```

Both point to same underlying array.

---

# 10. Deep Copy

```go
a := []int{1,2,3}

b := append([]int(nil),a...)
```

or

```go
b := make([]int,len(a))
copy(b,a)
```

---

# 11. Iterate Arrays/Slices

Index + Value

```go
for i,v := range nums{
    fmt.Println(i,v)
}
```

Only value

```go
for _,v := range nums{
    fmt.Println(v)
}
```

Only index

```go
for i := range nums{
    fmt.Println(i)
}
```

Classic

```go
for i:=0;i<len(nums);i++{
    fmt.Println(nums[i])
}
```

---

# 12. Delete Element

Delete index i

```go
nums = append(nums[:i], nums[i+1:]...)
```

Example

```
1 2 3 4

delete 2

↓

1 2 4
```

---

# 13. Insert Element

Insert x at i

```go
nums = append(nums[:i], append([]int{x}, nums[i:]...)...)
```

---

# 14. Reverse Slice

```go
for l,r:=0,len(nums)-1; l<r; l,r=l+1,r-1{
    nums[l],nums[r]=nums[r],nums[l]
}
```

---

# 15. Compare Arrays

Arrays

```go
a := [3]int{1,2,3}
b := [3]int{1,2,3}

fmt.Println(a==b)
```

Output

```
true
```

Slices

```go
a := []int{1,2,3}
b := []int{1,2,3}

fmt.Println(a==b)
```

❌ Not allowed (except comparing with `nil`).

---

# 16. Nil vs Empty Slice

```go
var a []int
```

```
nil
len=0
cap=0
```

---

```go
b := []int{}
```

```
not nil
len=0
cap=0
```

---

# 17. Passing to Functions

```go
func change(nums []int){
    nums[0]=100
}
```

```go
a:=[]int{1,2,3}

change(a)

fmt.Println(a)
```

Output

```
[100 2 3]
```

Because slices share underlying array.

---

# 18. Arrays are Copied

```go
func change(arr [3]int){
    arr[0]=100
}

a:=[3]int{1,2,3}

change(a)

fmt.Println(a)
```

Output

```
[1 2 3]
```

Whole array is copied.

---

# 19. Useful Built-in Functions

## len

```go
len(nums)
```

Returns number of elements.

---

## cap

```go
cap(nums)
```

Returns capacity.

---

## append

```go
nums = append(nums,10)
```

---

## copy

```go
copy(dst,src)
```

---

## make

```go
make([]int,5)
```

---

## new

```go
p := new([5]int)
```

Returns pointer.

---

# 20. Useful Packages

## sort

```go
sort.Ints(nums)
```

---

Descending

```go
sort.Sort(sort.Reverse(sort.IntSlice(nums)))
```

---

Binary Search

```go
idx := sort.SearchInts(nums,7)
```

---

Custom Sort

```go
sort.Slice(arr, func(i,j int) bool{
    return arr[i] < arr[j]
})
```

---

## slices (Go 1.21+)

```go
import "slices"
```

Contains

```go
slices.Sort(nums)

slices.Reverse(nums)

slices.Contains(nums,10)

slices.Index(nums,10)

slices.Equal(a,b)

slices.Clone(nums)

slices.Delete(nums,2,4)

slices.Insert(nums,2,100)

slices.BinarySearch(nums,8)

slices.Max(nums)

slices.Min(nums)
```

Very useful in interviews.

---

# 21. Interview Tricks

### Remove Last Element

```go
nums = nums[:len(nums)-1]
```

---

### Get Last Element

```go
last := nums[len(nums)-1]
```

---

### Queue

Push

```go
q = append(q,x)
```

Pop

```go
q = q[1:]
```

Front

```go
q[0]
```

---

### Stack

Push

```go
st = append(st,x)
```

Top

```go
st[len(st)-1]
```

Pop

```go
st = st[:len(st)-1]
```

---

### Clone Slice

```go
clone := append([]int(nil), nums...)
```

---

### Merge Two Slices

```go
a = append(a,b...)
```

---

### Check Empty

```go
if len(nums)==0{

}
```

---

# 22. Common Mistakes

❌ Forgetting to assign append result

```go
append(nums,5)
```

✅

```go
nums = append(nums,5)
```

---

❌ Comparing slices

```go
a==b
```

---

❌ Passing slice without `...`

```go
sum(nums)
```

---

✅

```go
sum(nums...)
```

---

❌ Assuming slice copy is deep copy

```go
b:=a
```

Both share same array.

---

# 23. Quick Revision (30 Seconds)

✅ Arrays
- Fixed size
- Value type
- Comparable
- Copied on assignment

✅ Slices
- Dynamic
- Backed by array
- Share memory
- Not comparable

✅ `...`
- In parameter → variadic
- In function call → expand slice

✅ Frequently Used
- `append`
- `copy`
- `len`
- `cap`
- `make`
- `sort`
- `slices`

---

# One-Line Memory Trick

```
Array = Fixed Box 📦

Slice = Window 🪟 into an Array

append() = May create a bigger box

copy() = Makes a new box

... in function = Accept many

... in call = Expand slice
```