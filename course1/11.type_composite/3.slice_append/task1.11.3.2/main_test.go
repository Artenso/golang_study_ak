package main

import (
	"reflect"
	"testing"
)

func TestAppendInt(t *testing.T) {
	type args struct {
		xs *[]int
		x  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "#1", args: args{xs: &[]int{1, 2, 3}, x: []int{4, 5, 6}}, want: []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if appendInt(tt.args.xs, tt.args.x...); !reflect.DeepEqual(*tt.args.xs, tt.want) {
				t.Errorf("got %v, want %v", *tt.args.xs, tt.want)
			}
		})
	}
}
