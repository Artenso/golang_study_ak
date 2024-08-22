package main

import "fmt"

func Pop(xs []int) (int, []int) {
	if len(xs) == 0 {
		return 0, []int{}
	}
	res := make([]int, len(xs)-1)
	copy(res, xs[1:])
	return xs[0], res
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	elem, res := Pop(xs)
	fmt.Printf("Значение: %d, Новый срез: %v\n", elem, res)
}
