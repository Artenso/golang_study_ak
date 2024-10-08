package main

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{name: "#1", args: args{xs: []int{1, 2, 3, 4, 5}}, want: 1, want1: []int{5, 1, 2, 3, 4}},
		{name: "#2", args: args{xs: []int{10, 9, 8, 7, 6, 5, 4}}, want: 10, want1: []int{4, 10, 9, 8, 7, 6, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Shift(tt.args.xs)
			if got != tt.want {
				t.Errorf("Shift() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Shift() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
