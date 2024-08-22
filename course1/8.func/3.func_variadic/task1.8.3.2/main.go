package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ConcatenateStrings("", "hello", "world", "how", "are", "you")) // even: hello-how-you, odd: world-are
}

func ConcatenateStrings(sep string, str ...string) string {
	even := make([]string, 0, len(str)/2)
	odd := make([]string, 0, len(str)/2)
	for i, str := range str {
		if (i & 1) == 0 {
			even = append(even, str)
			continue
		}
		odd = append(odd, str)
	}
	return fmt.Sprintf("even: %s, odd: %s", strings.Join(even, sep), strings.Join(odd, sep))
}
