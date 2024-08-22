package main

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	type args struct {
		xs []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{name: "#1", args: args{xs: []int{1, 2, 3, 4, 5}}, want: 1, want1: []int{2, 3, 4, 5}},
		{name: "#2", args: args{xs: []int{}}, want: 0, want1: []int{}},
		{name: "#3", args: args{xs: []int{10}}, want: 10, want1: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Pop(tt.args.xs)
			if got != tt.want {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
