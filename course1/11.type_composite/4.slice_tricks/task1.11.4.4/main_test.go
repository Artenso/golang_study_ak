package main

import (
	"reflect"
	"testing"
)

func TestRemoveIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "#1", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 2}, want: []int{1, 2, 4, 5}},
		{name: "#2", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 0}, want: []int{2, 3, 4, 5}},
		{name: "#3", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 15}, want: []int{}},
		{name: "#4", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 4}, want: []int{1, 2, 3, 4}},
		{name: "#5", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 5}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveIDX(tt.args.xs, tt.args.idx)
			// change element in base slise to check resolt slice isolation
			tt.args.xs[len(tt.args.xs)-1] = 666
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
