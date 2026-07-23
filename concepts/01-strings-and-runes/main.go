package main

import (
	"fmt"
	"strings"
)

func main() {
	showBytesAndRunes()
	showUnicodeSafeReverse()
	showBuilder()
}

func showBytesAndRunes() {
	text := "Go😊"

	fmt.Println("text:", text)
	fmt.Println("len in bytes:", len(text))
	fmt.Println("rune count:", len([]rune(text)))

	for byteIndex, r := range text {
		fmt.Printf("byte index %d: %c\n", byteIndex, r)
	}
}

func showUnicodeSafeReverse() {
	text := "Go😊"
	runes := []rune(text)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}

	fmt.Println("reversed by rune:", string(runes))
}

func showBuilder() {
	words := []string{"Go", "is", "fast"}
	var builder strings.Builder
	for i, word := range words {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}

	fmt.Println("built string:", builder.String())
}
