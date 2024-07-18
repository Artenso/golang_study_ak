package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	fileDir := filepath.Join(".", filepath.Dir(filePath))
	fileName := filepath.Base(filePath)

	err := os.MkdirAll(fileDir, perm)
	if err != nil {
		return fmt.Errorf("failed to create dir: %s", err.Error())
	}

	err = os.Chdir(fileDir)
	if err != nil {
		return fmt.Errorf("failed to move to created dir: %s", err.Error())
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, perm)
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err.Error())
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %s", err.Error())
	}

	return nil
}
func main() {
	err := WriteFile("/path/to/file.txt", []byte("Hello, World!"), os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
}
