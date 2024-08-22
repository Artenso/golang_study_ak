package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(CompareRoundedValues(2.5693, 2.5721, 3))
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	ratio := math.Pow(10, float64(decimalPlaces))
	ra := math.Round(a*ratio) / ratio
	rb := math.Round(b*ratio) / ratio
	if ra == rb {
		return true, 0.0
	}
	return false, math.Abs(ra - rb)
}
