package main

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func WithUsername(userName string) UserOption {
	return func(u *User) {
		u.Username = userName
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}

func NewUser(id int, opts ...UserOption) *User {
	user := &User{
		ID: id,
	}

	for _, opt := range opts {
		opt(user)
	}

	return user
}

func main() {
	user := NewUser(
		1,
		WithUsername("testuser"),
		WithEmail("testuser@example.com"),
		WithRole("admin"),
	)
	fmt.Printf("User: %+v\n", user)
}
