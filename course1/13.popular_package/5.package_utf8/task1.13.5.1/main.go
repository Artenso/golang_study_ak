package main

import (
	"fmt"
	"unicode/utf8"
)

func countUniqueUTF8Chars(s string) int {
	unicRunes := make(map[rune]struct{})
	offset := 0
	res := 0
	for offset < len(s)-1 {
		r, size := utf8.DecodeRuneInString(s[offset:])
		fmt.Println(string(r))
		offset += size
		if _, ok := unicRunes[r]; ok {
			continue
		}
		unicRunes[r] = struct{}{}
		res++
	}
	return res
}
