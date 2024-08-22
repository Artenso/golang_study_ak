package main

import (
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
)

func Test_getUniqueUsers(t *testing.T) {

	user1 := User{
		Nickname: gofakeit.FirstName(),
		Email:    gofakeit.Email(),
		Age:      gofakeit.Number(18, 60),
	}

	user2 := User{
		Nickname: gofakeit.FirstName(),
		Email:    gofakeit.Email(),
		Age:      gofakeit.Number(18, 60),
	}

	type args struct {
		users []User
	}
	tests := []struct {
		name  string
		args  args
		want  []User
		want1 int
	}{
		{name: "#1", args: args{users: []User{user1, user2, user2, user1}}, want: []User{user1, user2}, want1: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUniqueUsers(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUniqueUsers() = %v, want %v", got, tt.want)
			}
			if capacity := cap(getUniqueUsers(tt.args.users)); capacity != tt.want1 {
				t.Errorf("capacity = %v, want %v", capacity, tt.want)
			}
		})
	}
}
