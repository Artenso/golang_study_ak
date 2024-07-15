package main

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	unicUsers := make(map[string]User, len(users))
	for _, user := range users {
		if _, ok := unicUsers[user.Nickname]; ok {
			continue
		}
		unicUsers[user.Nickname] = user
	}

	res := make([]User, 0, len(unicUsers))
	for _, user := range unicUsers {
		res = append(res, user)
	}

	return res
}
