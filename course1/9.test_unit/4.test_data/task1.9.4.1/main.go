package main

import "math/rand"

func average(xs []float64) float64 {
	res := 0.0
	for _, num := range xs {
		res += num
	}
	return res / float64(len(xs))
}

func generateSlice(size int) []float64 {
	testData := make([]float64, 0, size)
	for i := 0; i < size; i++ {
		testData = append(testData, rand.Float64())
	}
	return testData
}
