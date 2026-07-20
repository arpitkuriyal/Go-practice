# Interfaces: Revision

## Mental model

An interface value contains a dynamic type and a dynamic value. An interface is `nil` only when both are nil.

```go
var p *User = nil
var x interface{} = p
fmt.Println(x == nil) // false: dynamic type is *User
```

## Rules to remember

| Topic | Rule |
| --- | --- |
| Satisfaction | Types implement interfaces implicitly by having the required method set. |
| Value receiver | Both `T` and `*T` have the method. |
| Pointer receiver | Only `*T` has the method. |
| Assertion | Use `value, ok := x.(T)` unless a failed assertion is impossible. |
| Type switch | Use for behaviour that depends on a small set of dynamic types. |
| Comparison | Interface values are comparable only when their dynamic values are comparable; slices/maps/functions cause a panic if compared. |

## Design guidance

- Accept interfaces where behaviour is needed; return concrete types unless you have a strong reason not to.
- Keep interfaces small. Define them near the consumer, not in a central “interfaces” package.
- Avoid `any` when a meaningful interface expresses the operation you need.
- Be careful calling methods on a typed nil stored inside an interface; the method may dereference the nil receiver.

## Interview answer: pointer or value receiver?

Use a pointer receiver when the method mutates the receiver, copying it would be expensive, or the type contains synchronization primitives. Prefer one receiver style consistently for a type.

Run the examples:

```bash
go run ./concepts/05-interfaces
```
