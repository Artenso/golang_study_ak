package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckDiscount(500, 70))
}

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return 0.0, fmt.Errorf("Cкидка не может превышать 50%")
	}
	return price - price*discount/100, nil
}
