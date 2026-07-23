# String Pattern Matching

## Challenge

Solve two related search problems:

1. Find every exact occurrence of a pattern, including overlaps.
2. Find every position where a permutation (anagram) of a pattern begins.

## Exact matching with overlaps

`FindAll` checks each possible starting position:

```go
for i := 0; i <= len(text)-len(pattern); i++ {
	if text[i:i+len(pattern)] == pattern {
		matches = append(matches, i)
	}
}
```

Advancing by one after each match includes overlaps:

```text
text:    aaaa
pattern: aa
matches: 0, 1, 2
```

It returns **byte indexes**, which is appropriate for byte-level exact matching but must be documented for Unicode input.

## Anagram matching: sliding window

`FindAnagramIndices` converts input to `[]rune` and maintains frequency maps for the wanted pattern and the current window.

```text
add next rune → remove rune leaving the window → compare counts
```

The window size always equals the pattern length. When both frequency maps match, the current window is an anagram.

## Complexity

| Function | Time | Index type |
| --- | --- | --- |
| `FindAll` | `O((n-m+1) × m)` | Byte index |
| `FindAnagramIndices` | `O(n × k)` with map comparison | Rune index |

`n` is text length, `m` is pattern length, and `k` is the number of distinct runes in a window. More advanced variants avoid a full map comparison by tracking a mismatch count.

## Edge cases

- Empty pattern returns no matches by this API decision.
- Pattern longer than text returns no matches.
- Exact matches can overlap.
- `FindAll` is byte-oriented; `FindAnagramIndices` is rune-oriented and reports rune positions.
- A rune is not always a user-perceived grapheme cluster; that distinction matters for advanced Unicode UI text.

## Interview answer

“For exact matching I scan each valid starting position and advance by one to keep overlaps. For anagrams I use a fixed-size sliding window and frequency counts. I state whether indexes are bytes or runes before implementing Unicode-sensitive behavior.”

## Test

```bash
go test ./challenges/08-string-pattern-matching
```
