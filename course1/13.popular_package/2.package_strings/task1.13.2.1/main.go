package main

import (
	"strings"
)

// func main() {
// 	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit
// amet ipsum mauris.
// Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.
// Praesent et diam eget libero egestas mattis sit amet vitae augue.`
// 	words := []string{"sit", "amet", "lorem"}
// 	result := CountWordsInText(txt, words)
// 	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
// }

func CountWordsInText(txt string, kyeWords []string) map[string]int {
	words := strings.Fields(txt)
	res := make(map[string]int, len(kyeWords))
	for _, kyeWord := range kyeWords {
		res[kyeWord] = 0
	}
	for _, word := range words {
		if _, ok := res[strings.ToLower(word)]; ok {
			res[strings.ToLower(word)]++
		}
	}
	return res
}
