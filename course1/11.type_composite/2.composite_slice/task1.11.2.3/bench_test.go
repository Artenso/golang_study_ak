package main

import "testing"

var testData = []int{1, 1, 2, 3, 4, 4, 3, 2, 5, 9, 4, 3, 9}

func BenchmarkFindSingleNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findSingleNumber(testData)
	}
}
