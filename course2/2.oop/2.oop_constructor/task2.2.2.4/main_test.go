package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestNewLogSystem(t *testing.T) {
	type args struct {
		opts []LogOption
	}

	fileLogger := &FileLogger{}
	consoleLogger := &ConsoleLogger{}

	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			name: "With file logger",
			args: args{
				opts: []LogOption{
					WithLogger(fileLogger),
				},
			},
			want: &LogSystem{
				logger: fileLogger,
			},
		},
		{
			name: "With console logger",
			args: args{
				opts: []LogOption{
					WithLogger(&ConsoleLogger{}),
				},
			},
			want: &LogSystem{
				logger: consoleLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogSystem(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogSystem_Log(t *testing.T) {
	type args struct {
		message string
	}

	file, err := os.Create("test_log.txt")
	if err != nil {
		log.Fatalf("failed to create test logfile: %s", err.Error())
	}
	defer os.RemoveAll("test_log.txt")
	defer file.Close()

	tests := []struct {
		name    string
		ls      *LogSystem
		args    args
		wantErr bool
	}{
		{
			name: "With file logger",
			ls: &LogSystem{
				logger: &FileLogger{
					file: file,
				},
			},
			args: args{
				message: "test log message",
			},
			wantErr: false,
		},
		{
			name: "With console logger",
			ls: &LogSystem{
				logger: &ConsoleLogger{
					out: os.Stdout,
				},
			},
			args: args{
				message: "test log message",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ls.Log(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("LogSystem.Log() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {

	expected := "Hello, world!"

	main()

	data, err := os.ReadFile("log.txt")
	if err != nil {
		log.Fatalf("failed to read log file: %s", err.Error())
	}

	if string(data) != expected {
		t.Errorf("got: %s, want: %s", string(data), expected)
	}
}
