package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func main() {
	jData := []byte(`{"name":"John","age":30}`)
	yData := []byte(`name: John
age: 30
`)
	var jPerson Person
	var yPerson Person

	if err := unmarshal(jData, &jPerson); err != nil {
		fmt.Println("Ошибка декодирования данных:", err)
		return
	}

	if err := unmarshal(yData, &yPerson); err != nil {
		fmt.Println("Ошибка декодирования данных:", err)
		return
	}
	fmt.Println("JSON:")
	fmt.Println("Имя:", jPerson.Name)
	fmt.Println("Возраст:", jPerson.Age)
	fmt.Println("YAML:")
	fmt.Println("Имя:", yPerson.Name)
	fmt.Println("Возраст:", yPerson.Age)
}

func unmarshal(data []byte, v interface{}) error {
	if json.Valid(data) {
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
		return nil
	}

	if err := yaml.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
