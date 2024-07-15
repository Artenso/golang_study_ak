package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(CalculatePercentageChange("621", "15"))
}

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	floatInitialValue, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0.0, err
	}

	if floatInitialValue == 0 {
		return 0.0, fmt.Errorf("zero initial value")
	}

	floatFinalValue, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0.0, err
	}

	pch := 100 - floatFinalValue/floatInitialValue*100
	return pch, nil
}
