package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Logger interface {
	Log(message string) error
}

type RemoteLogger struct {
	Address string
}

func (rl *RemoteLogger) Log(message string) error {

	jsonStr := fmt.Sprintf(`"message": %s`, message)
	buf := bytes.NewBufferString(jsonStr)

	req, err := http.NewRequest("POST", rl.Address, buf)
	if err != nil {
		return err
	}

	cli := &http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to write log message, status: %s", resp.Status)
	}

	return nil
}

type ConsoleLogger struct {
	Writer *os.File
}

func (cl *ConsoleLogger) Log(message string) error {
	_, err := cl.Writer.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

type FileLogger struct {
	File *os.File
}

func (fl *FileLogger) Log(message string) error {
	_, err := fl.File.WriteString(message)
	if err != nil {
		return err
	}
	return nil
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}
func main() {
	file, err := os.Create("course2/2.oop/1.oop_go/task2.2.1.2/loggs.log")
	if err != nil {
		log.Fatalf("failed to create logfile: %s", err.Error())
	}
	defer file.Close()

	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	fileLogger := &FileLogger{File: file} // Здесь замени на открытие реального файла, но для примера мы будем использовать os.Stdout
	loggers := []Logger{consoleLogger, fileLogger}
	LogAll(loggers, "This is a test log message.")
}
