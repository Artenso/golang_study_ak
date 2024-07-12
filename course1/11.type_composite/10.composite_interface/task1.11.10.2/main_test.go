package main

import (
	"reflect"
	"testing"
)

func TestConcat(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "#1", args: args{xs: []interface{}{"Hello ", "World!"}}, want: "Hello World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args struct {
		xs []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "#1", args: args{xs: []interface{}{1, 2, 3, 4, 5}}, want: 15},
		{name: "#2", args: args{xs: []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}}, want: 16.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.xs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperate(t *testing.T) {
	type args struct {
		f func(xs ...interface{}) interface{}
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{name: "#1", args: args{f: Concat, i: []interface{}{"Hello ", "World!"}}, want: "Hello World!"},
		{name: "#2", args: args{f: Sum, i: []interface{}{1, 2, 3, 4, 5}}, want: 15},
		{name: "#3", args: args{f: Sum, i: []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}}, want: 16.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Operate(tt.args.f, tt.args.i...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operate() = %v, want %v", got, tt.want)
			}
		})
	}
}
