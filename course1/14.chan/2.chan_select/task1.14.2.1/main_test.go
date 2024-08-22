package main

import "testing"

func Test_trySend(t *testing.T) {
	type args struct {
		ch chan int
		v  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1", args: args{ch: make(chan int, 2), v: 123}, want: true},
		{name: "#2", args: args{ch: make(chan int), v: 123}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trySend(tt.args.ch, tt.args.v)
			if got := trySend(tt.args.ch, tt.args.v); got != tt.want {
				t.Errorf("trySend() = %v, want %v", got, tt.want)
			}
		})
	}
}
