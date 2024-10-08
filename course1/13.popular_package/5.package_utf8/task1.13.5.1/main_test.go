package main

import (
	"testing"
)

func Test_countUniqueUTF8Chars(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#1", args: args{s: "Hello, world!"}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUniqueUTF8Chars(tt.args.s); got != tt.want {
				t.Errorf("countUniqueUTF8Chars() = %v, want %v", got, tt.want)
			}
		})
	}
}
