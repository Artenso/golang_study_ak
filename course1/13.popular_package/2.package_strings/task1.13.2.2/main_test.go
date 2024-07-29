package main

import (
	"reflect"
	"testing"
)

var user1 = User{
	Name: "Betty",
	Comments: []Comment{
		{Message: "good Comment 1"},
		{Message: "BaD CoMmEnT 2"},
		{Message: "Bad Comment 3"},
		{Message: "Use camelCase please"},
	},
}

var user2 = User{
	Name: "Jhon",
	Comments: []Comment{
		{Message: "Good Comment 1"},
		{Message: "Good Comment 2"},
		{Message: "Good Comment 3"},
		{Message: "Bad Comments 4"},
	},
}

var filteredUser1 = User{
	Name: "Betty",
	Comments: []Comment{
		{Message: "good Comment 1"},
		{Message: "Use camelCase please"},
	},
}

var filteredUser2 = User{
	Name: "Jhon",
	Comments: []Comment{
		{Message: "Good Comment 1"},
		{Message: "Good Comment 2"},
		{Message: "Good Comment 3"},
	},
}

func TestFilterComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{name: "#1", args: args{users: []User{user1, user2}}, want: []User{filteredUser1, filteredUser2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBadComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "#1", args: args{comment: "Good Comment 1"}, want: false},
		{name: "#2", args: args{comment: "BaD CoMmEnT 2"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadComment(tt.args.comment); got != tt.want {
				t.Errorf("IsBadComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBadComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{
			name: "#1",
			args: args{users: []User{user1, user2}},
			want: []Comment{{Message: "BaD CoMmEnT 2"}, {Message: "Bad Comment 3"}, {Message: "Bad Comments 4"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBadComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoodComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{
			name: "#1",
			args: args{users: []User{user1, user2}},
			want: []Comment{
				{Message: "good Comment 1"},
				{Message: "Use camelCase please"},
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoodComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoodComments() = %v, want %v", got, tt.want)
			}
		})
	}
}
