# Interfaces: Revision

## `any` vs `interface{}`

Since Go 1.18, `any` is simply an alias for `interface{}`.

```go
type any = interface{}
```

There is **no difference** in memory, performance, or behavior.

Use:

- `any` → when a value can be **any type**.
- `interface { ... }` → when defining **behavior** (methods).

```go
func Print(v any) {}
```

```go
type Speaker interface {
	Speak()
}
```

---

## Mental model

An interface value stores two things:

- **Dynamic type**
- **Dynamic value**

An interface is `nil` **only when both are nil**.

```go
var p *User = nil
var x any = p

fmt.Println(x == nil) // false
```

Internally:

```
Dynamic type  : *User
Dynamic value : nil
```

Since the dynamic type is `*User`, the interface itself is **not nil**.

---

## Rules to remember

| Topic | Rule |
| --- | --- |
| `any` vs `interface{}` | `any` is an alias for `interface{}`. Use `any` for arbitrary values and `interface` for defining behavior. |
| Interface nil | An interface is nil only when both its dynamic type and dynamic value are nil. |
| Interface satisfaction | Types implement interfaces **implicitly** by having the required method set. |
| Value receiver | Both `T` and `*T` implement the interface. |
| Pointer receiver | Only `*T` implements the interface. |
| Type assertion | Prefer `value, ok := x.(T)` to avoid panics. |
| Type switch | Use when behavior depends on a small set of dynamic types. |
| Comparison | Interface values are comparable only if their dynamic values are comparable. Slices, maps, and functions are not comparable (except with `nil`). |

---

## Design guidance

- Accept interfaces when your function depends on **behavior**, not a concrete type.
- Return concrete types unless an interface is genuinely required.
- Keep interfaces small (often one or two methods).
- Define interfaces close to where they are used, not in a central `interfaces` package.
- Use `any` only when a function truly accepts **any value**.
- If you only need one capability (for example, `Read()` or `Write()`), define a small interface instead of using `any`.
- Be careful calling methods on a typed `nil` stored inside an interface; the method may dereference the nil receiver.

---

## Interview answers

### Pointer receiver or value receiver?

Use a **pointer receiver** when:

- The method modifies the receiver.
- Copying the value would be expensive.
- The type contains synchronization primitives (`sync.Mutex`, `sync.RWMutex`, etc.).
- You want all methods on the type to use pointer receivers consistently.

Use a **value receiver** when:

- The method does not modify the receiver.
- The type is small (for example, `time.Time`).
- Value semantics make sense.

---

### Why is a nil pointer inside an interface not nil?

Because an interface stores both a **dynamic type** and a **dynamic value**.

```go
var p *User = nil
var x any = p

fmt.Println(x == nil) // false
```

The interface contains:

```
Dynamic type  : *User
Dynamic value : nil
```

Since the dynamic type exists, the interface itself is not `nil`.

---

### Why does this panic?

```go
var a any = []int{1}
var b any = []int{1}

fmt.Println(a == b)
```

Because the dynamic values are slices, and slices are **not comparable**.

---

### Why use `value, ok := x.(T)`?

A failed single-value type assertion panics.

```go
v := x.(string) // panic if x is not a string
```

The safe form avoids the panic.

```go
v, ok := x.(string)

if ok {
	fmt.Println(v)
}
```

---

## Run the examples

```bash
go run ./concepts/05-interfaces
```