package main

import "testing"

func TestCountBytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#1", args: args{str: "Привет, мир!"}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBytes(tt.args.str); got != tt.want {
				t.Errorf("CountBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSymbols(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#1", args: args{str: "Привет, мир!"}, want: 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSymbols(tt.args.str); got != tt.want {
				t.Errorf("CountSymbols() = %v, want %v", got, tt.want)
			}
		})
	}
}
