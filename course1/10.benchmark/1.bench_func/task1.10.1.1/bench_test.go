package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Fibonacci(5)
	}
}
