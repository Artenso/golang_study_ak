package main

import (
	"io"
	"log"
	"os"
)

func ReadString(filePath string) string {
	// open file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err.Error())
	}
	// close file in the end
	defer file.Close()

	// read file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %s", err.Error())
	}

	return string(data)
}
