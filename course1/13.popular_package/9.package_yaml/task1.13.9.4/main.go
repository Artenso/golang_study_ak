package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {
	// Write YAML data to file
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
	err := writeYAML("users.yaml", data)
	if err != nil {
		panic(err)
	}
}

func writeYAML(filePath string, data interface{}) error {
	dir := filepath.Dir(filePath)

	if err := os.MkdirAll(dir, os.FileMode(0644)); err != nil {
		return fmt.Errorf("failed to create dir: %s", err.Error())
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("failed to open file: %s", err.Error())
	}
	defer file.Close()

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal yaml: %s", err.Error())
	}

	_, err = file.Write(yamlData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %s", err.Error())
	}

	return nil
}
