# 02. Maps: Look Up Values by Key

A map stores a value under a key. Think of a phone book: use a name (key) to find a number (value).

```go
ages := map[string]int{"Arpit": 23}
ages["Neha"] = 24

fmt.Println(ages["Arpit"]) // 23
```

## Create a map

```go
counts := make(map[string]int)
settings := map[string]string{"theme": "dark"}
```

Use `make` when you want an empty map. Use a literal when you already know initial values.

## Read, write, delete

```go
counts["go"]++              // missing int key starts at 0
delete(counts, "old-value") // safe even if key is absent
```

Reading a missing key returns the zero value. Use the comma-ok form when zero could also be a real value:

```go
value, found := counts["go"]
if !found {
	// choose a default or handle absence
}
```

## Nil maps

```go
var counts map[string]int
fmt.Println(counts["go"]) // safe: 0
// counts["go"] = 1      // panic: map is nil
```

Initialize a map before writing to it.

## Maps share their stored data

Passing a map to a function lets that function change its entries:

```go
func addCount(counts map[string]int) {
	counts["go"]++
}
```

But assigning a new map inside that function does not replace the caller’s map variable.

## Rules to remember

- Map keys must be comparable: strings, numbers, booleans, pointers, and comparable structs work. Slices and maps do not.
- Map iteration order is deliberately unspecified. Sort keys when output order matters.
- Map values are not addressable. To update a stored struct, read it, change the copy, then store it back.
- A normal map is not safe for concurrent read/write access. Learn the safe options in the next map lesson.

## Interview answer

“A map gives fast key-to-value lookup. Missing keys return the zero value, so I use comma-ok when I must tell missing apart from zero. I initialize before writing, do not rely on iteration order, and synchronize shared maps.”

## Run

```bash
go run ./concepts/02-maps
```
