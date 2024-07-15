package main

import (
	"reflect"
	"testing"
)

func TestFilterDividers(t *testing.T) {
	type args struct {
		xs      []int
		divider int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "#1", args: args{xs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, divider: 3}, want: []int{3, 6, 9}},
		{name: "#2", args: args{xs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, divider: 0}, want: []int{}},
		{name: "#3", args: args{xs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, divider: -2}, want: []int{2, 4, 6, 8}},
		{name: "#4", args: args{xs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, divider: 50}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterDividers(tt.args.xs, tt.args.divider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterDividers() = %v, want %v", got, tt.want)
			}
		})
	}
}
