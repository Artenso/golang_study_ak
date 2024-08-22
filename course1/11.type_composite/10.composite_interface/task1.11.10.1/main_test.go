package main

import "testing"

func TestGetType(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{val: 1}, want: "int"},
		{name: "#2", args: args{val: 0.5}, want: "float64"},
		{name: "#3", args: args{val: "Hello World!"}, want: "string"},
		{name: "#4", args: args{val: interface{}(nil)}, want: "Empty interface"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.val); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
