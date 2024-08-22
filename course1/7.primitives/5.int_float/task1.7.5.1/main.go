package main

import (
	"fmt"
	"log"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000")) // 0.15625
}

func binaryStringToFloat(binary string) float32 {
	var number uint32
	// Преобразование строки в двоичной системе в целочисленное представление
	num64, err := strconv.ParseUint(binary, 2, 32)
	if err != nil {
		log.Fatalf("failed to parse binary string: %s", err.Error())
	}
	number = uint32(num64)
	// Преобразование целочисленного представления в число с плавающей точкой
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}
