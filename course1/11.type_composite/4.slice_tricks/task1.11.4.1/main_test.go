package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	type args struct {
		xs    []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ok", args: args{xs: []int{1, 2, 3, 4, 5}, start: 1, end: 3}, want: []int{2, 3, 4}},
		{name: "upBorder", args: args{xs: []int{1, 2, 3, 4, 5}, start: 1, end: 4}, want: []int{2, 3, 4, 5}},
		{name: "outOfBorder", args: args{xs: []int{1, 2, 3, 4, 5}, start: -2, end: 10}, want: []int{}},
		{name: "wrongBorders", args: args{xs: []int{1, 2, 3, 4, 5}, start: 3, end: 1}, want: []int{2, 3, 4}},
		{name: "end=start", args: args{xs: []int{1, 2, 3, 4, 5}, start: 2, end: 2}, want: []int{3}},
		{name: "all", args: args{xs: []int{1, 2, 3, 4, 5}, start: 0, end: 4}, want: []int{1, 2, 3, 4, 5}},
		{name: "empty", args: args{xs: []int{}, start: 0, end: 0}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cut(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cut() = %v, want %v", got, tt.want)
			}
		})
	}
}
