package main

import "testing"

func Test_getVariableType(t *testing.T) {
	type args struct {
		variable interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{variable: 10}, want: "int"},
		{name: "#2", args: args{variable: "hello"}, want: "string"},
		{name: "#3", args: args{variable: 0.256}, want: "float64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVariableType(tt.args.variable); got != tt.want {
				t.Errorf("getVariableType() = %v, want %v", got, tt.want)
			}
		})
	}
}
