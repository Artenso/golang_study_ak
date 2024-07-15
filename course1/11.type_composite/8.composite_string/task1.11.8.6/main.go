package main

import (
	"regexp"
)

func CountVowels(str string) int {
	reg := regexp.MustCompile(`[aeiouyаеёиоуыэюя]`)
	return len(reg.FindAllString(str, -1))
}
