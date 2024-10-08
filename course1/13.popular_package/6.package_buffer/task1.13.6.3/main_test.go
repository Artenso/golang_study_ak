package main

import (
	"bytes"
	"testing"
)

func Test_getDataString(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "#1", args: args{b: bytes.NewBufferString("Hello, World!")}, want: "Hello, World!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDataString(tt.args.b); got != tt.want {
				t.Errorf("getDataString() = %v, want %v", got, tt.want)
			}
		})
	}
}
