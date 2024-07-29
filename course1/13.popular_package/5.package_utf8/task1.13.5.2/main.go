package main

import (
	"strings"
	"unicode"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range strings.ToLower(s) {
		if isRussianLetter(char) {
			counts[char]++
		}
	}
	return counts
}

func isRussianLetter(char rune) bool {
	return unicode.Is(unicode.Cyrillic, char)
}
