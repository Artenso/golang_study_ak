package main

import "fmt"

func main() {
	var name, city string
	var age int

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите ваш возаст: ")
	fmt.Scan(&age)
	fmt.Print("Введите ваш город: ")
	fmt.Scan(&city)
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age)
	fmt.Println("Город:", city)
}
