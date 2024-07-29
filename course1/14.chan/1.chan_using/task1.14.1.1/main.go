package main

import (
	"fmt"
	"sync"
)

func mergeChan(mergeTo chan int, from ...chan int) {
	defer close(mergeTo)
	wg := sync.WaitGroup{}

	for _, in := range from {
		wg.Add(1)
		go func(in chan int) {
			defer wg.Done()
			for data := range in {
				mergeTo <- data
			}

		}(in)
	}

	wg.Wait()
}

func mergeChan2(chans ...chan int) chan int {
	wg := sync.WaitGroup{}

	mergeTo := make(chan int, 16)

	for _, in := range chans {
		wg.Add(1)
		go func(in chan int) {
			defer wg.Done()
			for data := range in {
				mergeTo <- data
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(mergeTo)
	}()

	return mergeTo
}

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch1 := generateChan(10)
	ch2 := generateChan(15)
	ch3 := generateChan(20)

	mergeTo := mergeChan2(ch1, ch2, ch3)

	for data := range mergeTo {
		fmt.Println(data)
	}

	fmt.Println("====================================")

	mergeTo = make(chan int)

	ch1 = generateChan(10)
	ch2 = generateChan(15)
	ch3 = generateChan(20)

	go mergeChan(mergeTo, ch1, ch2, ch3)

	for data := range mergeTo {
		fmt.Println(data)
	}
}
