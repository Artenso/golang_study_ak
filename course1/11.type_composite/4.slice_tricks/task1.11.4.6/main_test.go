package main

import (
	"reflect"
	"testing"
)

func TestInsertToStart(t *testing.T) {
	type args struct {
		xs []int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "#1", args: args{xs: []int{4, 5, 6}, x: []int{1, 2, 3}}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "#2", args: args{xs: []int{4, 5, 6}}, want: []int{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := tt.args.xs
			got := InsertToStart(tt.args.xs, tt.args.x...)
			// check base slice isn't changed
			if !reflect.DeepEqual(before, tt.args.xs) {
				t.Errorf("base slice before = %v, after %v", before, tt.args.xs)
			}
			// change element in base slise to check resolt slice isolation
			tt.args.xs[len(tt.args.xs)-1] = 666
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIDX() = %v, want %v", got, tt.want)
			}
		})
	}
}
