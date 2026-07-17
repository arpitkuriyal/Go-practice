package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordfrequencyCounter(str string) map[string]int {
	seen := make(map[string]int)
	var cleanStr strings.Builder
	for _, ch := range str {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == ' ' {
			cleanStr.WriteRune(unicode.ToLower(ch))
		}
	}
	s := strings.SplitSeq(cleanStr.String(), " ")
	for v := range s {
		seen[v]++
	}
	return seen
}
func main() {
	ans := wordfrequencyCounter("hello.... world# hello123 world hello h h h h")
	fmt.Println(ans)
}
