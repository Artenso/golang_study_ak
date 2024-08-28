package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() map[string]bool {
	res := make(map[string]bool, 2)

	start := time.Now()
	_ = factorialRecursive(100000)
	recursiveTime := time.Since(start)

	start = time.Now()
	_ = factorialIterative(100000)
	iterativeTime := time.Since(start)

	if recursiveTime > iterativeTime {
		res["iterative"] = true
		res["recursive"] = false
	} else {
		res["iterative"] = false
		res["recursive"] = true
	}

	return res
}

func main() {
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH)
	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}

// Go version: go1.22.0
// Go OS/Arch: linux / amd64
// Which factorial is faster?
// map[iterative:true recursive:false]
