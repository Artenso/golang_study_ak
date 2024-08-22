package main

import (
	"fmt"
	"strings"
)

func main() {
	v1 := "0"
	v2 := "1"
	v3 := "1"
	v4 := "2"
	v5 := "3"
	v6 := "5"
	v7 := "8"
	v8 := "13"

	str := strings.Join([]string{v1, v2, v3, v4, v5, v6, v7, v8}, " ")
	fmt.Println(str)
}
