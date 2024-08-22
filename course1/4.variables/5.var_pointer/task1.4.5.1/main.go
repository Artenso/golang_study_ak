package main

import "fmt"

func main() {
	num := 5
	flnum := 2.53
	str := "Hello World!"
	boo := true

	a := &num
	b := &flnum
	c := &str
	d := &boo

	changeInt(a)
	changeFloat(b)
	changeString(c)
	changeBool(d)

	fmt.Println(num)
	fmt.Println(flnum)
	fmt.Println(str)
	fmt.Println(boo)
}

// меняет значение переменной на 20
func changeInt(a *int) {
	*a = 20
}

// меняет значение переменной на 6.28
func changeFloat(b *float64) {
	*b = 6.28
}

// меняет значение переменной на "Goodbye, world!"
func changeString(c *string) {
	*c = "Goodbye, world!"
}

// меняет значение переменной на false
func changeBool(d *bool) {
	*d = false
}
