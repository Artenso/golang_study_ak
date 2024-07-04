package main

import "fmt"

func main() {
	num := 5
	str := "Hello World!"

	mutate(&num)
	ReverseString(&str)

	fmt.Println(num)
	fmt.Println(str)
}

func mutate(a *int) {
	*a = 42
}

func ReverseString(str *string) {
	r := []rune(*str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	*str = string(r)
}
