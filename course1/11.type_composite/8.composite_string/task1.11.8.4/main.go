package main

import "strings"

// fast
func concatStringsJoin(xs ...string) string {
	return strings.Join(xs, "")
}

// slow
// i was trying to make it faster and it becomes,
// but not enough, and it's become too complicate
func concatStrings(xs ...string) string {
	builder := strings.Builder{}
	n := 0
	for _, s := range xs {
		n += len([]byte(s))
	}
	builder.Grow(n)

	for _, s := range xs {
		builder.WriteString(s)
	}

	return builder.String()
}
