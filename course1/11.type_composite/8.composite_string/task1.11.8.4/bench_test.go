package main

import "testing"

func BenchmarkConcatStrings(b *testing.B) {
	xs := []string{"Hello", " ", "World", "!"}
	for i := 0; i < b.N; i++ {
		_ = concatStrings(xs...)
	}
}

func BenchmarkConcatStringsJoin(b *testing.B) {
	xs := []string{"Hello", " ", "World", "!"}
	for i := 0; i < b.N; i++ {
		_ = concatStringsJoin(xs...)
	}
}
