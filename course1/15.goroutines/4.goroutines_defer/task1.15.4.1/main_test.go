package main

import (
	"log"
	"os"
	"testing"
)

func Test_writeToFile(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "#1", args: args{data: "Hello, World!"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, err := os.Create("test.txt")
			defer os.RemoveAll(file.Name())
			if err != nil {
				log.Fatalf("failed to create file: %s", err.Error())
			}
			if err := writeToFile(file, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
