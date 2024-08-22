package main

import "fmt"

func GetType(val interface{}) string {
	switch val.(type) {
	case nil:
		return "Empty interface"
	default:
		return fmt.Sprintf("%T", val)
	}
}
