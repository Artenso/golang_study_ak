package main

import (
	"testing"
)

var testData = make([]int, 0, 7)

func TestRemoveExtraMemory(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#1", args: args{xs: append(testData, 1, 2, 3, 4)}, want: 4},
		{name: "#2", want: 0},
		{name: "#3", args: args{xs: append(testData, 1, 2, 3, 4, 5, 6, 7, 8, 9)}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cap(RemoveExtraMemory(tt.args.xs)); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
