package main

import "fmt"

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))  // Output: 5
	fmt.Println(MathOperate(Mul, 1, 7, 3))  // Output: 21
	fmt.Println(MathOperate(Sub, 13, 2, 3)) // Output: 8
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func Sum(a ...int) int {
	res := 0
	for _, n := range a {
		res += n
	}
	return res
}
func Mul(a ...int) int {
	res := 1
	for _, n := range a {
		res *= n
	}
	return res
}
func Sub(a ...int) int {
	var res int
	for i, n := range a {
		if i == 0 {
			res = n
			continue
		}
		res -= n
	}
	return res
}
