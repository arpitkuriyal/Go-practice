# Go Strings🚀


---

# 1. What is a String?

A string is a **read-only sequence of bytes**.

```go
s := "hello"
```

Internally

```
+---+---+---+---+---+
| h | e | l | l | o |
+---+---+---+---+---+
 104 101 108 108 111
```

Each element is a **byte (uint8)**.

---

# 2. Why are Strings Immutable?

You **cannot change** a character inside a string.

❌

```go
s := "hello"

s[0] = 'H'
```

Compiler error

```
cannot assign to s[0]
```

---

## Why?

A string is stored as

```
String

↓

Pointer → bytes
Length
```

```
        Pointer
           │
           ▼

+---+---+---+---+---+
| h | e | l | l | o |
+---+---+---+---+---+
```

The bytes are **read-only**.

This gives Go:

- Faster copying
- Memory sharing
- Safety
- Better performance

---

## Think Like This

```
Array

Editable

[1 2 3]
 ↑

Can modify


String

Read Only

"hello"

Cannot modify
```

---

# 3. Then How Do We Modify a String?

Convert it.

## Option 1 — []byte

```go
s := "hello"

b := []byte(s)

b[0] = 'H'

s = string(b)

fmt.Println(s)
```

Output

```
Hello
```

---

## Option 2 — []rune (Unicode Safe)

```go
s := "नमस्ते"

r := []rune(s)

r[0] = 'क'

fmt.Println(string(r))
```

---

# 4. String Length

```go
s := "hello"

fmt.Println(len(s))
```

Output

```
5
```

---

## Important

`len()` returns **bytes**, not characters.

Example

```go
s := "😊"

fmt.Println(len(s))
```

Output

```
4
```

Although you see

```
😊
```

it occupies **4 bytes** in UTF-8.

---

# 5. Indexing

```go
s := "hello"

fmt.Println(s[0])
```

Output

```
104
```

Why?

Because

```
'h'

↓

ASCII

104
```

Type

```go
byte
```

---

## Convert to Character

```go
fmt.Printf("%c\n", s[0])
```

Output

```
h
```

---

# 6. What Does s[i] Return?

Always

```
byte
```

NOT

```
string
```

Example

```go
s := "Go"

fmt.Println(s[1])
```

Output

```
111
```

Type

```go
byte
```

---

# 7. Iterate Over String

## Method 1

```go
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
```

Output

```
104
101
108
108
111
```

Works byte-by-byte.

---

## Method 2 (Most Common)

```go
for i, ch := range s {
    fmt.Println(i, ch)
}
```

Output

```
0 104
1 101
2 108
3 108
4 111
```

For ASCII it looks similar.

---

# 8. Unicode Example

```go
s := "नमस्ते"
```

Using

```go
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
```

You'll get

```
224
164
168
...
```

because UTF-8 characters use multiple bytes.

---

Using

```go
for i, ch := range s {
    fmt.Printf("%d %c\n", i, ch)
}
```

Output

```
0 न
3 म
6 स
9 ्
12 त
15 े
```

Notice

```
0
3
6
9
```

because each rune uses multiple bytes.

---

# 9. byte vs rune

## byte

```
uint8

1 byte
```

Example

```
'A'

↓

65
```

---

## rune

```
int32

Unicode code point
```

Example

```
'😊'
```

needs a rune.

---

Rule

```
ASCII

↓

byte


Unicode

↓

rune
```

---

# 10. Convert String

String → Bytes

```go
b := []byte(s)
```

---

Bytes → String

```go
s := string(b)
```

---

String → Rune

```go
r := []rune(s)
```

---

Rune → String

```go
s := string(r)
```

---

# 11. String Concatenation

```go
a := "Hello"

b := "World"

fmt.Println(a + " " + b)
```

Output

```
Hello World
```

---

# 12. Compare Strings

```go
fmt.Println("abc"=="abc")
```

Output

```
true
```

---

Lexicographical

```go
fmt.Println("abc" < "abd")
```

Output

```
true
```

---

# 13. Strings Package

```go
import "strings"
```

---

Contains

```go
strings.Contains()

strings.HasPrefix()

strings.HasSuffix()

strings.Split()

strings.Join()

strings.ReplaceAll()

strings.TrimSpace()

strings.Trim()

strings.ToUpper()

strings.ToLower()

strings.Index()

strings.LastIndex()

strings.Repeat()

strings.Fields()

strings.Count()
```

---

Examples

```go
strings.Contains("golang","go")
```

true

---

```go
strings.Split("a,b,c",",")
```

↓

```
["a","b","c"]
```

---

```go
strings.Join(arr,",")
```

↓

```
a,b,c
```

---

```go
strings.ReplaceAll("go go","go","Go")
```

↓

```
Go Go
```

---

# 14. Building Large Strings

Bad

```go
s := ""

for i:=0;i<1000;i++{
    s += "Go"
}
```

Every `+` creates a **new string**, making this inefficient.

---

Good

```go
var b strings.Builder

b.WriteString("Go")
b.WriteString("Lang")

fmt.Println(b.String())
```

Use `strings.Builder` for repeated concatenation.

---

# 15. Reverse String

ASCII

```go
b := []byte(s)

for l,r:=0,len(b)-1;l<r;l,r=l+1,r-1{
    b[l],b[r]=b[r],b[l]
}

fmt.Println(string(b))
```

---

Unicode Safe

```go
r := []rune(s)

for l,rp:=0,len(r)-1;l<rp;l,rp=l+1,rp-1{
    r[l],r[rp]=r[rp],r[l]
}

fmt.Println(string(r))
```

---

# 16. Substring

```go
s := "golang"

fmt.Println(s[1:4])
```

Output

```
ola
```

---

# 17. String Formatting

```go
fmt.Sprintf()
```

Example

```go
name := "Go"

s := fmt.Sprintf("Hello %s", name)
```

---

# 18. Escape Characters

```go
"\n"
```

New line

---

```go
"\t"
```

Tab

---

```go
"\""
```

Double quote

---

```go
"\\"
```

Backslash

---

# 19. Raw String

```go
s := `Hello
World`
```

Output

```
Hello
World
```

No escaping needed.

---

# 20. Useful Functions

Length

```go
len(s)
```

---

Compare

```go
strings.Compare(a,b)
```

---

Equal Fold

```go
strings.EqualFold(a,b)
```

Case-insensitive comparison.

---

Repeat

```go
strings.Repeat("Go",3)
```

↓

```
GoGoGo
```

---

Trim

```go
strings.TrimSpace(s)
```

---

# 21. Common Interview Questions

## Reverse String

Use

```
[]rune
```

---

## Is Palindrome

Compare from both ends.

---

## Valid Anagram

Count frequencies (map or array).

---

## Longest Common Prefix

Compare characters column by column.

---

## Count Characters

```go
map[rune]int
```

---

## First Unique Character

Frequency map + second pass.

---

# 22. Common Mistakes

❌ Trying to modify string

```go
s[0]='H'
```

---

❌ Using byte loop for Unicode

```go
for i:=0;i<len(s);i++
```

on

```
😊
नमस्ते
こんにちは
```

---

✅ Use

```go
for _, r := range s
```

or

```go
[]rune(s)
```

---

❌ Assuming len() gives characters

```go
len("😊")
```

returns

```
4
```

bytes, not 1 character.

---

# 23. Quick Revision (30 Seconds)

✅ String = immutable sequence of bytes

✅ `s[i]` returns a `byte`, not a string

✅ `len(s)` returns bytes, not characters

✅ Use `range` or `[]rune` for Unicode

✅ Convert to `[]byte` or `[]rune` to modify

✅ Use `strings.Builder` for efficient concatenation

✅ Use the `strings` package for common operations

---

# One-Line Memory Trick

```
String = Read-only bytes 📜

byte = 1 byte (ASCII)

rune = Unicode character 🌍

range = Decodes runes

[]byte = Editable bytes

[]rune = Editable Unicode

strings.Builder = Fast string construction 🚀
```
