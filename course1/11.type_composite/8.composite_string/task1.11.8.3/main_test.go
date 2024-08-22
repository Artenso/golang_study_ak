package main

import (
	"reflect"
	"testing"
)

func Test_getBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "#1", args: args{s: "Hello World!"}, want: []byte("Hello World!")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRunes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: "#1", args: args{s: "Hello World!"}, want: []rune("Hello World!")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRunes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}
