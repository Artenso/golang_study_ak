package main

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit"
)

type User struct {
	Name string
	Age  int
}

// func main() {
// 	users := getUsers()
// 	res := preparePrint(users)

// 	fmt.Println(res)
// }

func GetUsers() []User {
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		users[i] = User{
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(18, 60),
		}
	}
	return users
}

func PreparePrint(users []User) string {
	str := make([]string, len(users))
	for i, user := range users {
		str[i] = fmt.Sprintf("Имя:  %s,  Возраст:  %d", user.Name, user.Age)
	}
	return strings.Join(str, "\n")
}
