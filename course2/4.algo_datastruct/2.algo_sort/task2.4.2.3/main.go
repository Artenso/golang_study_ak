package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/brianvoe/gofakeit"
)

// Структура пользователя
type User struct {
	ID   int
	Name string
	Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) == 0 {
		return arr2
	}

	if len(arr2) == 0 {
		return arr1
	}

	res := make([]User, 0, len(arr1)+len(arr2))

	i := 0
	j := 0

	for {
		if i < len(arr1) && j < len(arr2) {
			if arr1[i].ID <= arr2[j].ID {
				res = append(res, arr1[i])
				i++
				continue
			} else {
				res = append(res, arr2[j])
				j++
				continue
			}
		}

		if i >= len(arr1) && j >= len(arr2) {
			break
		}

		if i >= len(arr1) {
			res = append(res, arr2[j:]...)
			break
		}

		if j >= len(arr2) {
			res = append(res, arr1[i:]...)
			break
		}

	}

	return res
}

func generateUsers(n int) []User {
	gofakeit.Seed(time.Now().UnixNano())

	users := make([]User, n)
	for i := range users {
		users[i] = User{
			ID:   gofakeit.Number(0, 100),
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(18, 70),
		}
	}

	return users
}

func main() {
	users1 := generateUsers(7)
	users2 := generateUsers(3)
	sort.Slice(users1, func(i, j int) bool { return users1[i].ID < users1[j].ID })
	sort.Slice(users2, func(i, j int) bool { return users2[i].ID < users2[j].ID })
	fmt.Println("Первый список:\n", users1)
	fmt.Println("\nВторой список:\n", users2)
	fmt.Println("\nОбщий список:\n", Merge(users1, users2))
}
