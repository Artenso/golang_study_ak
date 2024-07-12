package main

import (
	"reflect"
	"testing"
)

func Test_getSubSlice(t *testing.T) {
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
		{name: "1", args: args{xs: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, start: 2, end: 6}, want: []int{3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubSlice(tt.args.xs, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSubSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
