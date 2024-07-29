package main

import (
	"os"
	"testing"
)

func TestReadString(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{filePath: "./test.txt"},
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.OpenFile(tt.args.filePath, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				t.Errorf("failed to create and open test file: %s", err.Error())
			}
			defer os.Remove(tt.args.filePath)

			file.WriteString(tt.want)
			file.Close()

			if got := ReadString(tt.args.filePath); got != tt.want {
				t.Errorf("ReadString() = %v, want %v", got, tt.want)
			}
		})
	}
}
