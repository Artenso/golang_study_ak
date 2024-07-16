package main

import (
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

// func main() {
// 	users := []User{
// 		{
// 			Name: "Betty",
// 			Comments: []Comment{
// 				{Message: "good Comment 1"},
// 				{Message: "BaD CoMmEnT 2"},
// 				{Message: "Bad Comment 3"},
// 				{Message: "Use camelCase please"},
// 			},
// 		},
// 		{
// 			Name: "Jhon",
// 			Comments: []Comment{
// 				{Message: "Good Comment 1"},
// 				{Message: "Good Comment 2"},
// 				{Message: "Good Comment 3"},
// 				{Message: "Bad Comments 4"},
// 			},
// 		},
// 	}
// 	users = FilterComments(users)
// 	fmt.Println(users)

// }

func FilterComments(users []User) []User {
	res := make([]User, 0, len(users))
	for _, user := range users {
		goodUser := User{Name: user.Name}
		for _, comment := range user.Comments {
			if IsBadComment(comment.Message) {
				continue
			}
			goodUser.Comments = append(goodUser.Comments, comment)
		}
		res = append(res, goodUser)
	}
	return res
}

func IsBadComment(comment string) bool {
	return strings.Contains(strings.ToLower(comment), "bad")
}

func GetBadComments(users []User) []Comment {
	badComments := make([]Comment, 0, 8)
	for _, user := range users {
		for _, comment := range user.Comments {
			if IsBadComment(comment.Message) {
				badComments = append(badComments, comment)
			}
		}

	}
	return badComments
}

func GetGoodComments(users []User) []Comment {
	goodComments := make([]Comment, 0, 8)
	for _, user := range users {
		for _, comment := range user.Comments {
			if IsBadComment(comment.Message) {
				continue
			}
			goodComments = append(goodComments, comment)
		}
	}
	return goodComments
}
