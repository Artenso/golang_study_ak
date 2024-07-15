package main

import (
	"strings"
)

func createUniqueText(text string) string {
	words := strings.Split(text, " ")
	check := make(map[string]struct{})
	res := make([]string, 0, len(words))
	for _, word := range words {
		if _, ok := check[word]; ok {
			continue
		}
		res = append(res, word)
		check[word] = struct{}{}
	}
	return strings.Join(res, " ")
}
