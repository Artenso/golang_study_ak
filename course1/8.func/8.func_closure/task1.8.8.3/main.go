package main

import "fmt"

func main() {
	// пример использования функции adder
	addTwo := adder(2)
	result := addTwo(3)
	result1 := addTwo(5)
	fmt.Println(result)  // выводит 5
	fmt.Println(result1) // 7
}

func adder(initial int) func(int) int {
	return func(i int) int {
		return initial + i
	}
}
