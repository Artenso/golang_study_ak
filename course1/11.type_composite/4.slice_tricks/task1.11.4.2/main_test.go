package main

import (
	"reflect"
	"testing"
)

func TestInsertAfterIDX(t *testing.T) {
	type args struct {
		xs  []int
		idx int
		x   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ok", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 2, x: []int{6, 7, 8}}, want: []int{1, 2, 3, 6, 7, 8, 4, 5}},
		{name: "idx=0", args: args{xs: []int{1, 2, 3, 4, 5}, idx: 0, x: []int{6, 7, 8}}, want: []int{1, 6, 7, 8, 2, 3, 4, 5}},
		{name: "no_xs_idx=-1", args: args{idx: -1, x: []int{6, 7, 8}}, want: []int{6, 7, 8}},
		{name: "no_xs_idx=0", args: args{idx: 0, x: []int{6, 7, 8}}, want: []int{}},
		{name: "empty_xs_idx>len", args: args{xs: []int{}, idx: 2, x: []int{6, 7, 8}}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertAfterIDX(tt.args.xs, tt.args.idx, tt.args.x...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAfterIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
