package main

import (
	"strings"
)

// func main() {
// 	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
// 	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
// 	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
// }

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

func Concat(xs ...interface{}) interface{} {
	tmp := make([]string, 0, len(xs))
	for _, x := range xs {
		if str, ok := x.(string); ok {
			tmp = append(tmp, str)
		}
	}
	return strings.Join(tmp, "")
}

func Sum(xs ...interface{}) interface{} {
	ints := 0
	floats := 0.0
	for _, x := range xs {
		if num, ok := x.(int); ok {
			ints += num
			continue
		}
		if num, ok := x.(float64); ok {
			floats += num
		}
	}
	if floats != 0 {
		return floats
	}
	return ints
}
