package main

import "unicode/utf8"

func CountBytes(str string) int {
	return len([]byte(str))
}

func CountSymbols(str string) int {
	return utf8.RuneCountInString(str)
}
