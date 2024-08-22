package main

import (
	"fmt"
	"strings"
)

func main() {
	// Test case 1
	result := UserInfo("John", 21, "Moscow", "Saint Petersburg")
	fmt.Println(result) // Имя: John, возраст: 21, города: Moscow, Saint Petersburg
	// Test case 2
	result = UserInfo("Alex", 34)
	fmt.Println(result) // Имя: Alex, возраст: 34, города:
}

func UserInfo(name string, age int, cities ...string) string {
	return fmt.Sprintf(
		"Имя: %s, возраст: %d, города: %s",
		name, age, strings.Join(cities, ", "),
	)
}
