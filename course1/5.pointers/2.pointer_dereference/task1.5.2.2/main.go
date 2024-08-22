package main

import (
	"fmt"
)

func main() {
	num := 6 // 720
	str1 := "abcdcba"
	str2 := "hello world"
	nums := []int{2, 3, 5, 3, 6, 4, 5, 7, 8, 4, 2, 6, 6}
	fmt.Println(Factorial(&num))
	fmt.Println(isPalindrome(&str1))
	fmt.Println(CountOccurrences(&nums, &num))
	fmt.Println(ReverseString(&str2))
}

func Factorial(n *int) int {
	num := *n
	res := 1
	for num >= 1 {
		res *= num
		num--
	}

	return res
}

func isPalindrome(str *string) bool {
	rstr := []rune(*str)
	for l, r := 0, len(rstr)-1; l < len(rstr)/2; l, r = l+1, r-1 {
		if rstr[l] != rstr[r] {
			return false
		}
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	res := 0
	for _, num := range *numbers {
		if num == *target {
			res++
		}
	}
	return res
}

func ReverseString(str *string) string {
	r := []rune(*str)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
