// Package patternmatching contains byte- and rune-aware string search examples.
package patternmatching

// FindAll returns every byte index where pattern occurs in text. Overlapping matches
// are included. An empty pattern returns no matches.
func FindAll(text, pattern string) []int {
	if pattern == "" || len(pattern) > len(text) {
		return []int{}
	}

	matches := make([]int, 0)
	for i := 0; i <= len(text)-len(pattern); i++ {
		if text[i:i+len(pattern)] == pattern {
			matches = append(matches, i)
		}
	}
	return matches
}

// FindAnagramIndices returns rune indexes where a permutation of pattern begins.
// It is Unicode-safe because it works over []rune rather than bytes.
func FindAnagramIndices(text, pattern string) []int {
	textRunes, patternRunes := []rune(text), []rune(pattern)
	if len(patternRunes) == 0 || len(patternRunes) > len(textRunes) {
		return []int{}
	}

	wanted := make(map[rune]int, len(patternRunes))
	window := make(map[rune]int, len(patternRunes))
	for _, r := range patternRunes {
		wanted[r]++
	}

	matches := make([]int, 0)
	for i, r := range textRunes {
		window[r]++
		if i >= len(patternRunes) {
			outgoing := textRunes[i-len(patternRunes)]
			window[outgoing]--
			if window[outgoing] == 0 {
				delete(window, outgoing)
			}
		}
		if i >= len(patternRunes)-1 && sameCounts(window, wanted) {
			matches = append(matches, i-len(patternRunes)+1)
		}
	}
	return matches
}

func sameCounts(left, right map[rune]int) bool {
	if len(left) != len(right) {
		return false
	}
	for r, count := range left {
		if right[r] != count {
			return false
		}
	}
	return true
}
