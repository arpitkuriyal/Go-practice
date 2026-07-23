# Word Frequency

## Challenge

Return a map containing the normalized frequency of every word in a sentence.

```text
input:  "Go, go! Go 1"
output: map[go:3 1:1]
```

## Core idea

Normalize first, then count. The solution ranges over runes so Unicode letters and digits can be recognized with the `unicode` package.

```go
counts := make(map[string]int)
counts[word]++ // a missing key starts at int's zero value: 0
```

`strings.Builder` efficiently creates normalized text without repeatedly allocating a new string for every append.

## Why a map?

A map gives expected `O(1)` lookup and increment per word. The overall runtime is `O(n)` for an input of `n` characters, plus the cost of splitting into words.

## Normalization decisions

Before implementing, define what a “word” means:

- Convert case with `unicode.ToLower`.
- Keep letters and digits when that matches the problem.
- Decide whether punctuation separates words (`"go,lang"` → `go`, `lang`) or is removed (`golang`).
- Ignore empty tokens from repeated or leading spaces.

These rules are part of the API contract, not just implementation detail.

## Edge cases

- Empty or whitespace-only input should normally return an empty map.
- Multiple spaces should not create an empty-string key.
- Unicode text should be processed as runes, not byte-by-byte.
- Map iteration order is unspecified; sort keys before producing deterministic display output.

## Interview answer

“I normalize the input once, split it into valid words, and use a map counter. Missing map keys have the zero value, so `counts[word]++` is enough. I clarify punctuation and Unicode rules before coding.”

## Run

```bash
go run ./challenges/04-word-frequency/word-frequency.go
```
