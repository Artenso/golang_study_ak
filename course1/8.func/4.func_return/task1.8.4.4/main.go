package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CalculatePercentageChange(1200, 300))
}

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	pch := 100 - finalValue/initialValue*100
	return math.Round(pch*100) / 100
}
