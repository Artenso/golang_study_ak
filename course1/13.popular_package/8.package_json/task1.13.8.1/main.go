package main

import "encoding/json"

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Text string `json:"text"`
}

func getJSON(data []User) (string, error) {
	res, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
