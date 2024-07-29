package main

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func Test_getReader(t *testing.T) {
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want *bufio.Reader
	}{
		{
			name: "#1",
			args: args{b: bytes.NewBufferString("Hello, World!")},
			want: bufio.NewReader(bytes.NewBufferString("Hello, World!")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReader(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
