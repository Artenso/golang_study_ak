package main

import "testing"

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "no elems", in: []int{}, want: 0},
		{name: "one elem", in: []int{1}, want: 0},
		{name: "many elems", in: []int{1, 5, 0, 30, -7, 100}, want: 107},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDifference(tt.in); got != tt.want {
				t.Errorf("MaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
