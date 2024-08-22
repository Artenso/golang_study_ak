package main

import "fmt"

func main() {
	d, r := DivideAndRemainder(5, 10)
	fmt.Printf("Частное: %d, Остаток: %d", d, r)
}

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			// обработка ошибки деления на ноль
			fmt.Println("check zero argument")
		}
	}()

	return a / b, a % b
}
