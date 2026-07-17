# 📌 Section 2: Follow-up Questions (Important) — Go Slices & Memory

This section covers important follow-up concepts frequently asked in interviews, especially around **slices, memory behavior, and references in Go**.

---

## 🔹 Q1: Fix Append

```go
func appendFix(s *[]int) {
	*s = append(*s, 100)
}
```

### 💡 Explanation

* Slices are passed **by value**, meaning the slice descriptor (pointer, length, capacity) is copied.
* If `append()` causes reallocation, the new array is not reflected outside the function.
* By passing a **pointer to slice (`*[]int`)**, we modify the original slice itself.

👉 **Key Idea:** Use a pointer when you want changes (like append) to persist outside the function.

---

## 🔹 Q2: Avoid Reallocation

```go
s := make([]int, 0, 100)
```

### 💡 Explanation

* Creates a slice with:

  * Length = 0
  * Capacity = 100
* Appending elements within capacity does **not trigger reallocation**.

👉 **Why important?**

* Reduces memory allocations
* Improves performance in high-scale systems

---

## 🔹 Q3: Avoid Shared Memory

```go
b := append([]int{}, arr[:2]...)
```

### 💡 Explanation

* Creates a **new slice with its own underlying array**
* Prevents unintended side effects on the original slice

👉 **Key Idea:** This is a **deep copy pattern** for slices.

---

## 🔹 Q4: Why `b` Didn’t Change?

### 💡 Explanation

* When `append()` exceeds capacity:

  * Go allocates a **new underlying array**
* The original slice still points to the **old array**

👉 **Conclusion:**

* New slice → new array
* Old slice → unchanged

---

## 🔹 Q5: Access Beyond Length

```go
s := make([]int, 3, 5)
fmt.Println(s[3]) // ❌ panic

s = s[:4] // ✅ valid
```

### 💡 Explanation

* You can only access elements **within length**, not capacity
* Index `3` becomes valid only after extending the slice

👉 **Rule:**

* Capacity ≠ Accessible range
* Length defines valid indexing

---

## 🔹 Q6: Slice Internals

### 💡 Explanation

A slice in Go is a **descriptor**, not the actual data.

It contains:

1. Pointer → address of underlying array
2. Length → number of accessible elements
3. Capacity → total space available

👉 **Key Insight:** Copying a slice copies this descriptor, not the data.

---

## 🔹 Q7: Map vs Slice

### 💡 Explanation

| Feature         | Slice             | Map              |
| --------------- | ----------------- | ---------------- |
| Type            | Descriptor        | Reference type   |
| Copy Behavior   | Descriptor copied | Reference shared |
| Underlying Data | Array             | Hash table       |

👉 **Key Difference:**

* Slice → partial reference (descriptor)
* Map → full reference

---

## 🔹 Q8: Struct with Slice

```go
type Data struct {
	arr []int
}

func modify(d Data) {
	d.arr[0] = 100
}
```

### 💡 Explanation

* Struct is passed **by value**
* But slice inside it points to the **same underlying array**

👉 **Result:**

* Modifying `d.arr` affects original data

---

## 🔹 Q9: Deep Copy

```go
newSlice := append([]int{}, oldSlice...)
```

### 💡 Explanation

* Creates a completely **independent copy**
* No shared memory between slices

👉 **Use when:**

* You want safe modifications without affecting original data

---

## 🔹 Q10: Capacity-Based Behavior

```go
func tricky(s []int) {
	s = append(s, 1)
	s[0] = 999
}
```

### 💡 Explanation

* Behavior depends on capacity:

#### Case 1: No Reallocation

* Append fits within capacity
* Same underlying array used
* ✅ Original slice gets modified

#### Case 2: Reallocation Happens

* New array created
* ❌ Original slice remains unchanged

👉 **Key Insight:**
Always think: *“Will append reallocate?”*

---

# 🧠 Final Mental Model

* Slice = **(Pointer + Length + Capacity)**
* `append()` may:

  * Modify same array OR
  * Allocate new array
* Sharing vs copying depends on:

  * Capacity
  * How slice is created

---

# 🚀 Interview Tip

If stuck, always ask yourself:

1. Is this slice sharing memory?
2. Will `append()` trigger reallocation?
3. Am I modifying underlying array or just the descriptor?

👉 These 3 questions solve 90% of slice problems.
