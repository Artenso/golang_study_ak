package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Write JSON data to file
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}
	err := writeJSON("users.json", data)
	if err != nil {
		panic(err)
	}
}

func writeJSON(filePath string, data interface{}) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, os.FileMode(0644)); err != nil {
		return fmt.Errorf("failed to create dir: %s", err.Error())
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err.Error())
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %s", err.Error())
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %s", err.Error())
	}

	return nil
}
