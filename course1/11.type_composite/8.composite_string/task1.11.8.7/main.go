package main

import "strings"

func ReplaceSymbols(str, old, new string) string {
	return strings.ReplaceAll(str, old, new)
}
