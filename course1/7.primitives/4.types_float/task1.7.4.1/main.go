package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypotenuse(5, 12))
}

func hypotenuse(a, b float64) float64 {
	return math.Hypot(a, b)
}
