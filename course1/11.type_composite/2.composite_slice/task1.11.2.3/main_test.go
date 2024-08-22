package main

import (
	"testing"
)

func TestBitwiseXOR(t *testing.T) {
	type args struct {
		n1 int
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "equal", args: args{n1: 123, n2: 123}, want: 0},
		{name: "diff", args: args{n1: 123, n2: 321}, want: 314},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseXOR(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("bitwiseXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want int
	}{
		{name: "empty", in: []int{}, want: 0},
		{name: "no single", in: []int{1, 1, 2, 3, 4, 4, 3, 2}, want: 0},
		{name: "with single", in: []int{1, 1, 2, 3, 4, 4, 3, 2, 5, 9, 4, 3, 9}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSingleNumber(tt.in); got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
