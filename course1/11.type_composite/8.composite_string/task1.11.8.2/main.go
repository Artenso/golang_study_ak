package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "Hello, World!"
	header := getStringHeader(s)
	fmt.Printf("Data: %v\n", header.Data)
	fmt.Printf("Len: %v\n", header.Len)
}

func getStringHeader(s string) reflect.StringHeader {
	return reflect.StringHeader{
		Data: *(*uintptr)(unsafe.Pointer(&s)),
		Len:  len(s),
	}
}
