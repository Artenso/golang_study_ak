package main

import (
	"io"
	"os"
)

// Logger interface
type Logger interface {
	Log(message string) error
}

// FileLogger struct
type FileLogger struct {
	file *os.File
}

func (fl *FileLogger) Log(message string) error {
	_, err := fl.file.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

// ConsoleLogger struct
type ConsoleLogger struct {
	out io.ReadWriter
}

func (cl *ConsoleLogger) Log(message string) error {
	_, err := cl.out.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

type LogSystem struct {
	logger Logger
}

// LogOption functional option type
type LogOption func(*LogSystem)

func WithLogger(logger Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = logger
	}
}

func NewLogSystem(opts ...LogOption) Logger {
	logger := &LogSystem{}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (ls *LogSystem) Log(message string) error {
	ls.logger.Log(message)
	return nil
}

func main() {
	file, _ := os.Create("course2/2.oop/2.oop_constructor/task2.2.2.4/log.txt")
	defer file.Close()
	fileLogger := &FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))
	logSystem.Log("Hello, world!")
}
