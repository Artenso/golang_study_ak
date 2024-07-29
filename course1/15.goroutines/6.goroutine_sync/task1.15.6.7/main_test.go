package main

import (
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
		// time.Sleep(1 * time.Microsecond)
		p = &Person{
			Age: gofakeit.Number(18, 60),
		}

		_ = p.Age
	}

}

func BenchmarkWithPool(b *testing.B) {
	gofakeit.Seed(0)
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	// benchmark code
	for i := 0; i < b.N; i++ {
		// time.Sleep(1 * time.Microsecond)
		p = personPool.Get().(*Person)

		p.Age = gofakeit.Number(18, 60)

		personPool.Put(p)
	}
}

/*

BenchmarkWithoutPool-4   	26053312	        47.26 ns/op	       8 B/op	       1 allocs/op
BenchmarkWithPool-4      	26043475	        43.66 ns/op	       0 B/op	       0 allocs/op

*/
