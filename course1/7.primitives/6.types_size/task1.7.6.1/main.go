package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b bool
	var n int

	fmt.Println("bool size:", sizeOfBool(b))
	fmt.Println("int size:", sizeOfInt(n))
	fmt.Println("int8 size:", sizeOfInt8(int8(n)))
	fmt.Println("int16 size:", sizeOfInt16(int16(n)))
	fmt.Println("int32 size:", sizeOfInt32(int32(n)))
	fmt.Println("int64 size:", sizeOfInt64(int64(n)))
	fmt.Println("uint size:", sizeOfUint(uint(n)))
	fmt.Println("uint8 size:", sizeOfUint8(uint8(n)))

}

func sizeOfBool(b bool) int {
	res := unsafe.Sizeof(b)
	return int(res)
}
func sizeOfInt(n int) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfInt8(n int8) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfInt16(n int16) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfInt32(n int32) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfInt64(n int64) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfUint(n uint) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
func sizeOfUint8(n uint8) int {
	res := unsafe.Sizeof(n)
	return int(res)
}
