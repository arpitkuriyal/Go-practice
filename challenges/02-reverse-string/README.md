# Reverse a String

## Challenge

Read a string and return it in reverse order.

```text
input:  Go
output: oG
```

The included solution uses a two-pointer swap. It is a good ASCII/byte-based solution. For user-facing text, reverse runes instead so UTF-8 characters are not corrupted.

## Core idea: two pointers

1. Convert the immutable string to a mutable slice.
2. Start one pointer at the beginning and one at the end.
3. Swap the values and move both pointers inward.
4. Convert the slice back to a string.

```go
b := []byte(s)
for left, right := 0, len(b)-1; left < right; left, right = left+1, right-1 {
	b[left], b[right] = b[right], b[left]
}
return string(b)
```

## Why convert first?

Strings are immutable in Go. You can read `s[0]`, but cannot assign to it:

```go
s := "hello"
// s[0] = 'H' // compile error
```

Converting to `[]byte` or `[]rune` creates a mutable sequence that can be swapped.

## `byte` versus `rune`

| Type | Meaning | Use it when |
| --- | --- | --- |
| `byte` | An 8-bit value; a UTF-8 byte | Working with ASCII, bytes, or protocol data. |
| `rune` | An `int32` Unicode code point | Working with human-readable Unicode text. |

`len(s)` and `s[i]` operate on bytes, not visible characters:

```go
fmt.Println(len("Go")) // 2
fmt.Println(len("😊")) // 4: UTF-8 bytes
```

For ASCII, a byte and a character line up. For Unicode text, they often do not.

## Unicode-safe version

The byte version reverses the individual UTF-8 bytes of `"नमस्ते"` or `"😊"`, producing invalid or incorrect text. Convert to runes for code-point-safe reversal:

```go
func reverseUnicode(s string) string {
	r := []rune(s)
	for left, right := 0, len(r)-1; left < right; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	return string(r)
}
```

`range` also decodes a string into runes:

```go
for byteIndex, r := range "Go😊" {
	fmt.Println(byteIndex, r)
}
```

The index from `range` is a **byte index**, while `r` is a rune.

## Important Unicode caveat

Runes are Unicode code points, not always user-perceived characters (grapheme clusters). For example, some emoji sequences and letters plus combining accents contain more than one rune. `[]rune` is the correct standard-library answer for most interview problems, but full grapheme-aware reversal requires a Unicode text-segmentation library and a clearly defined product requirement.

## Complexity

| Approach | Time | Extra space |
| --- | --- | --- |
| `[]byte` two pointers | `O(n)` bytes | `O(n)` |
| `[]rune` two pointers | `O(n)` runes | `O(n)` |

The extra space is needed because Go strings cannot be modified in place.

## Edge cases

- Empty string: returns `""`.
- One character: returns the same string.
- Palindrome: returns the same sequence.
- Unicode input: use the rune version.
- Spaces: `fmt.Scan` reads only up to whitespace; use `bufio.Reader` if the input can be a sentence.

## Interview answer

“Because strings are immutable, I convert the input to a mutable slice and swap from both ends in `O(n)` time. I use `[]byte` for ASCII and `[]rune` for Unicode text, because indexing a string and `len` operate on UTF-8 bytes.”

## Run

```bash
go run ./challenges/02-reverse-string/reverse.go
```
