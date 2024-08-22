package main

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "#1", want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(GetUsers()); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreparePrint(t *testing.T) {
	type args struct {
		users []User
	}

	users := []User{
		{Name: gofakeit.Name(), Age: gofakeit.Number(18, 60)},
		{Name: gofakeit.Name(), Age: gofakeit.Number(18, 60)},
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "#1",
			args: args{users: users},
			want: fmt.Sprintf(
				"Имя:  %s,  Возраст:  %d\nИмя:  %s,  Возраст:  %d",
				users[0].Name, users[0].Age,
				users[1].Name, users[1].Age,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreparePrint(tt.args.users); got != tt.want {
				t.Errorf("preparePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}
