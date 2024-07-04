package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func main() {
	fmt.Println(GenerateFakeData())
	fmt.Println(GenerateFakeData())
}

func GenerateFakeData() string {
	res := fmt.Sprintf(
		"Name: %s\nAddress: %s\nPhone: %s\nEmail: %s\n",
		fake.FullName(),
		fake.StreetAddress(),
		fake.Phone(),
		fake.EmailAddress(),
	)
	return res
}
