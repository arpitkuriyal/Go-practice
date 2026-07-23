# 01. Strings, Bytes, and Runes

## Start simple

A Go string is read-only text stored as bytes. For simple English text, one character usually uses one byte:

```go
text := "Go"
fmt.Println(len(text)) // 2 bytes
```

Strings cannot be changed in place:

```go
// text[0] = 'g' // compile error
```

Create a new string instead when you need a changed value.

## `byte` versus `rune`

UTF-8 is the common encoding for Go strings. Some characters use more than one byte:

```go
text := "😊"
fmt.Println(len(text)) // 4 bytes
```

| Type | Meaning | Use it for |
| --- | --- | --- |
| `byte` | One 8-bit UTF-8 byte. | Raw bytes, ASCII text, protocols. |
| `rune` | One Unicode code point (`int32`). | Human-readable Unicode text. |

`text[i]` returns a byte, not a character. Use `range` to read runes:

```go
for byteIndex, r := range "Go😊" {
	fmt.Println(byteIndex, r)
}
```

The index from `range` is still a byte index; the value `r` is a rune.

## When to use `[]byte` or `[]rune`

Use `[]byte` when working with ASCII or raw data:

```go
data := []byte("hello")
data[0] = 'H'
text := string(data) // "Hello"
```

Use `[]rune` when changing or reversing Unicode text:

```go
runes := []rune("Go😊")
// swap or change runes safely
text := string(runes)
```

## Build strings efficiently

Repeated `+` concatenation in a large loop can create many temporary strings. Use `strings.Builder` when you are gradually building text:

```go
var builder strings.Builder
for _, word := range words {
	builder.WriteString(word)
}
result := builder.String()
```

For a few strings, `first + " " + last` is perfectly clear and fine.

## Important caveat

A rune is a Unicode code point, but not always one user-perceived character. Some emoji and accented characters are made from several runes. `[]rune` is the right answer for most Go text problems; full grapheme handling needs a dedicated Unicode library and a clear product requirement.

## Interview answer

“Go strings are immutable UTF-8 byte sequences. `len` and indexing work in bytes, so I use `range` or `[]rune` for Unicode text. For large incremental string building, I use `strings.Builder`.”

## Run

```bash
go run ./concepts/01-strings-and-runes
```
