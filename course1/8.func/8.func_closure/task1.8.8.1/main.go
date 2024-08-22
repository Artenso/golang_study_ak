package main

import "fmt"

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3

	counter = createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
	fmt.Println(counter()) // 4
	fmt.Println(counter()) // 5
	fmt.Println(counter()) // 6
}

func createCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
