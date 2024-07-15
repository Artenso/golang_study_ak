package main

import "testing"

func TestCountVowels(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "#1", args: args{str: "Hello World!"}, want: 3},
		{name: "#2", args: args{str: "Привет мир"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountVowels(tt.args.str); got != tt.want {
				t.Errorf("CountVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
