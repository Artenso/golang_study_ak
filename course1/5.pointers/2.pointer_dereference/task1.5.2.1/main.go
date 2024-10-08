package main

import "fmt"

// Пример использования функций
func main() {
	a := 5
	b := 10
	c := Dereference(&a)
	d := Sum(&b, &c)
	fmt.Println(c) // Output: 5
	fmt.Println(d) // Output: 15
}

func Dereference(n *int) int {
	return *n
}

func Sum(a, b *int) int {
	return *a + *b
}
