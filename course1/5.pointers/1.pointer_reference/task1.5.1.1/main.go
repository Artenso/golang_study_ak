package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
)

func main() {
	numbers := []int{2, 3, 8, 5, 3, 4, 6, 1}
	strs := []string{"hello", "world"}

	fmt.Println(*Add(numbers[0], numbers[2]))
	fmt.Println(*Max(numbers))
	fmt.Println(*IsPrime(numbers[1]))
	fmt.Println(*ConcatenateStrings(strs))
}

func Add(a, b int) *int {
	res := a + b
	return &res
}

func Max(numbers []int) *int {
	sort.Ints(numbers)
	return &numbers[len(numbers)-1]
}

func IsPrime(number int) *bool {
	res := big.NewInt(int64(number)).ProbablyPrime(0)
	return &res
}

func ConcatenateStrings(strs []string) *string {
	res := strings.Join(strs, " ")
	return &res
}
