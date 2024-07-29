package main

import (
	"fmt"
	"os"
	"sync"
)

var ch = make(chan int)

const (
	one = iota + 1
	two
	three
	four
	five
)

func main() {
	numbers := []int{one, two, three, four, five}
	storeNumbers(numbers)
	print(ch)
}

func print(data chan int) {
	if len(os.Getenv("DEBUG")) != 0 {
		return
	}

	for v := range data {
		fmt.Println(v)
	}
}

func storeNumbers(numbers []int) {
	wg := sync.WaitGroup{}
	for _, num := range numbers {
		wg.Add(1)
		go func(num int) { // исправить, но не убирать анонимную функцию
			write(num)
			wg.Done()
		}(num)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
}

func write(n int) {
	ch <- n
}
