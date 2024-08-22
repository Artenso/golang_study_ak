package main

import (
	"testing"
)

func Test_concatStrings(t *testing.T) {
	type args struct {
		xs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{xs: []string{"Hello", " ", "World", "!"}}, want: "Hello World!"},
		{name: "#2", args: args{xs: []string{"Hello"}}, want: "Hello"},
		{name: "#3", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatStrings(tt.args.xs...); got != tt.want {
				t.Errorf("concatStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concatStringsJoin(t *testing.T) {
	type args struct {
		xs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{xs: []string{"Hello", " ", "World", "!"}}, want: "Hello World!"},
		{name: "#2", args: args{xs: []string{"Hello"}}, want: "Hello"},
		{name: "#3", want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatStringsJoin(tt.args.xs...); got != tt.want {
				t.Errorf("concatStringsJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
