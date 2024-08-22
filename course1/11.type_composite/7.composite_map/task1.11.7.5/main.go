package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
// 	filter := map[string]bool{"ipsum": true, "elit": true}
// 	filteredSentence := filterSentence(sentence, filter)
// 	fmt.Println(filteredSentence)
// }

// slower
func filterSentenceReplace(sentence string, filter map[string]bool) string {
	for bad := range filter {
		sentence = strings.ReplaceAll(sentence, fmt.Sprintf(" %s", bad), "")
		sentence = strings.ReplaceAll(sentence, fmt.Sprintf("%s ", bad), "")
	}
	return sentence
}

// faster
func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Split(sentence, " ")
	res := make([]string, 0, len(words))
	for _, word := range words {
		if _, ok := filter[word]; ok {
			continue
		}
		res = append(res, word)
	}
	return strings.Join(res, " ")
}
