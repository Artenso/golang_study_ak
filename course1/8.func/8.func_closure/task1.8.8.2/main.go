package main

import "fmt"

func main() {
	// Пример использования функции multiplier
	m := multiplier(2.5)
	result := m(100)
	fmt.Println(result) // Вывод: 250
}

func multiplier(factor float64) func(float64) float64 {
	return func(f float64) float64 {
		return factor * f
	}
}
