# String Pattern Matching

`FindAll` returns byte indexes for exact substring matches, including overlaps. `FindAnagramIndices` uses runes and a sliding window, so its indexes are Unicode-safe rune positions.

```bash
go test ./challenges/08-string-pattern-matching
```
