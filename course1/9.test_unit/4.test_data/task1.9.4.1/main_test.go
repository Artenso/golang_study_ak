package main

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	type testData struct {
		in  []float64
		exp float64
	}

	testSlice1 := []float64{1, 2, 3, 4, 5}
	testSlice2 := []float64{2, 1, 3, 5, 4}
	testSlice3 := []float64{2, 1, 3}

	testCases := []testData{
		{in: testSlice1, exp: 3},
		{in: testSlice2, exp: 3},
		{in: testSlice3, exp: 2},
	}

	for _, tc := range testCases {
		res := average(tc.in)
		if res != tc.exp {
			t.Errorf("average(%v) = %f; wanted: %f", tc.in, res, tc.exp)
		}
	}
}

func TestGenerateSlice(t *testing.T) {
	res1 := generateSlice(5)
	res2 := generateSlice(5)
	if reflect.DeepEqual(res1, res2) {
		t.Errorf("slices are equal")
	}
}
