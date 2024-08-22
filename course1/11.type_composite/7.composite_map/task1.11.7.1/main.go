package main

import (
	"regexp"
	"strings"
)

// func main() {
// 	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
// 	occurrences := countWordOccurrences(text)
// 	for word, count := range occurrences {
// 		fmt.Printf("%s: %d\n", word, count)
// 	}
// }

func countWordOccurrences(text string) map[string]int {
	text = strings.ToLower(text)

	reg := regexp.MustCompile(`[[:punct:]]`)
	text = reg.ReplaceAllString(text, "")

	words := strings.Split(text, " ")
	counts := make(map[string]int, len(words))
	for _, word := range words {
		counts[word]++
	}
	return counts
}
