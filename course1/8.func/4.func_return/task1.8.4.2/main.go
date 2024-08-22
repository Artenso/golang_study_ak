package main

import (
	"fmt"
)

func main() {
	fmt.Println(FindMaxAndMin(3, 4, 3, 90, 0, 435, 3, 5, 7, 8, 9, 1, 3, 56, 324, 12, -54))
}

func FindMaxAndMin(n ...int) (int, int) {
	if len(n) < 1 {
		fmt.Println("no arguments")
		return 0, 0
	}
	min := n[0]
	max := n[0]
	for _, num := range n {
		if num > max {
			max = num
			continue
		}
		if num < min {
			min = num
		}

	}
	return max, min
}
