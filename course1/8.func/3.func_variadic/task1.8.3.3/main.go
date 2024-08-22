package main

import "fmt"

func main() {
	PrintNumbers(1, 2, 3, 4, 5)
}

func PrintNumbers(numbers ...int) {
	for _, num := range numbers {
		fmt.Println(num)
	}
}
