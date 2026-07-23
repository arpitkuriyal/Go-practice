# 04. Structs, Methods, and Pointers

## Start simple: group related data

A struct creates your own type with named fields:

```go
type User struct {
	Name string
	Age  int
}

user := User{Name: "Arpit", Age: 23}
fmt.Println(user.Name)
```

Use a struct when several values describe one thing.

## Methods: behavior belonging to a type

A method is a function with a receiver before its name:

```go
func (u User) Greeting() string {
	return "Hello, " + u.Name
}
```

Call it with a value:

```go
fmt.Println(user.Greeting())
```

## Value receiver versus pointer receiver

This method receives a copy of the user, so changing it does not change the caller:

```go
func (u User) Rename(name string) {
	u.Name = name // changes only the copy
}
```

Use a pointer receiver when a method must change the original value:

```go
func (u *User) Birthday() {
	u.Age++
}
```

Go lets you call `user.Birthday()` and automatically takes the address when `user` is addressable.

## When should I choose each?

| Choose a value receiver when | Choose a pointer receiver when |
| --- | --- |
| The method only reads a small value. | The method changes the original value. |
| Copying the value is cheap. | Copying would be expensive. |
| Value semantics are useful. | The type contains a mutex or other non-copyable state. |

Keep receiver style consistent for one type. If one important method needs a pointer receiver, using pointer receivers for its other methods is often clearest.

## Pointers in one minute

A pointer stores an address. `&user` means “the address of user.” `*pointer` means “the value at that address.”

```go
user := User{Name: "Arpit"}
pointer := &user
pointer.Name = "Neha" // shorthand for (*pointer).Name
```

The original `user` now has the name `Neha`.

## Embedding

Embedding places one type inside another:

```go
type Employee struct {
	User
	Role string
}
```

An `Employee` can use `employee.Name` because `Name` comes from its embedded `User`. Embedding is composition—a way to reuse fields and methods—not inheritance.

## Rules to remember

- Struct values are copied when passed or assigned.
- Use `==` only when every field in the struct is comparable.
- Prefer named fields in literals when a struct has several fields: `User{Name: "Arpit", Age: 23}`.
- Do not copy a struct containing a `sync.Mutex` after using it.

## Interview answer

“Structs group related data. Methods give behavior to a type. I use value receivers for small read-only values and pointer receivers when a method mutates the value, copying is expensive, or the type contains synchronization state. Embedding is composition, not inheritance.”

## Run

```bash
go run ./concepts/04-structs-methods-and-pointers
```
