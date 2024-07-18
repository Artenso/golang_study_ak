package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/brianvoe/gofakeit"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var DB = &sql.DB{}

func initDbConn() {
	db, err := sql.Open("sqlite3", "./course1/13.popular_package/10.package_db/task1.13.10.1/users.db")
	if err != nil {
		log.Fatalf("failed to connect to db: %s", err.Error())
	}
	DB = db
}

func CreateUserTable() error {
	_, err := DB.Exec(`CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
		)`)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(user User) error {
	_, err := DB.Exec("INSERT INTO users (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(id int) (User, error) {
	rows, err := DB.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user := User{}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func UpdateUser(user User) error {
	_, err := DB.Exec("UPDATE users SET age = ?, name = ? WHERE id = ?", user.Age, user.Name, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	initDbConn()
	defer DB.Close()

	if err := CreateUserTable(); err != nil {
		fmt.Println(err)
	}

	user := User{
		Name: gofakeit.Name(),
		Age:  gofakeit.Number(18, 60),
	}

	if err := InsertUser(user); err != nil {
		fmt.Println(err)
	}

	dbUser, err := SelectUser(9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbUser)

	dbUser.Name = gofakeit.Name()

	if err := UpdateUser(dbUser); err != nil {
		fmt.Println(err)
	}

	dbUser, err = SelectUser(dbUser.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbUser)

	if err := DeleteUser(9); err != nil {
		fmt.Println(err)
	}

	dbUser, err = SelectUser(9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dbUser)
}
