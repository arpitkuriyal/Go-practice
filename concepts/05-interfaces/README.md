# 05. Interfaces: Work With Behavior, Not One Concrete Type

An interface describes methods a value must have. It answers: “what can this value do?”

```go
type Speaker interface {
	Speak()
}

type Dog struct{}
func (Dog) Speak() { fmt.Println("woof") }

func sayHello(s Speaker) {
	s.Speak()
}
```

`Dog` satisfies `Speaker` automatically because it has a `Speak` method. Go has no `implements` keyword.

## Why interfaces help

The `sayHello` function does not need to know whether it received a `Dog`, `Person`, or another type. It only needs something that can `Speak`.

Keep interfaces small and define them near the code that uses them.

```go
type Reader interface {
	Read([]byte) (int, error)
}
```

## `any` and `interface{}`

`any` is another name for `interface{}`:

```go
func printValue(value any) {}
```

Use `any` when a value can truly be any type. Use a named interface when your function needs a specific behavior.

## Value receiver versus pointer receiver

```go
type Cat struct{}
func (Cat) Speak() {}
```

Both `Cat{}` and `&Cat{}` satisfy `Speaker` because `Speak` has a value receiver.

```go
type Dog struct{}
func (*Dog) Speak() {}
```

Only `&Dog{}` satisfies `Speaker` because `Speak` has a pointer receiver. Use pointer receivers when a method changes the value or copying it would be expensive.

## Type assertions

When an interface may hold different concrete types, check safely:

```go
value, ok := input.(string)
if ok {
	fmt.Println(value)
}
```

The one-value form, `input.(string)`, panics when the value is not a string. Use the comma-ok form unless failure is impossible by design.

## The typed-nil surprise

```go
var user *User = nil
var value any = user

fmt.Println(value == nil) // false
```

The interface remembers that its value has type `*User`, even though the pointer inside it is nil. Check carefully before calling methods that may dereference that pointer.

## Rules to remember

- Return concrete types by default; accept an interface when you depend on behavior.
- Interfaces are satisfied implicitly.
- An interface is `nil` only when it has no stored type and no stored value.
- Interface comparison panics if the stored value is not comparable, such as a slice or map.

## Interview answer

“An interface is a small contract of behavior. Types satisfy it implicitly. I accept interfaces when a function needs a capability, keep them small, and use safe type assertions when the concrete type is not guaranteed.”

## Run

```bash
go run ./concepts/05-interfaces
```
