package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	animals := []Animal{
		{Type: "Cat", Name: "Barsik", Age: 2},
		{Type: "Dog", Name: "Reks", Age: 3},
		{Type: "Horse", Name: "Spirit", Age: 1},
	}
	return animals
}
func preparePrint(animals []Animal) string {
	str := make([]string, len(animals))
	for i, animal := range animals {
		str[i] = fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d", animal.Type, animal.Name, animal.Age)
	}
	return strings.Join(str, "\n")
}
