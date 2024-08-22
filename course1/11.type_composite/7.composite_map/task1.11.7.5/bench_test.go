package main

import "testing"

func BenchmarkFilterSentence(b *testing.B) {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}
	for i := 0; i < b.N; i++ {
		_ = filterSentence(sentence, filter)
	}
}

func BenchmarkFilterSentenceReplace(b *testing.B) {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}
	for i := 0; i < b.N; i++ {
		_ = filterSentenceReplace(sentence, filter)
	}
}
