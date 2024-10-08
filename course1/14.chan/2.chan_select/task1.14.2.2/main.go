package main

import (
	"fmt"
	"time"
)

func main() {
	timeoutFunc := timeout(4 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}

func timeout(timeout time.Duration) func() bool {
	ch := make(chan struct{})

	go func() {
		time.Sleep(timeout)
		ch <- struct{}{}
		close(ch)
	}()

	return func() bool {
		select {
		case <-ch:
			return true
		default:
			return false
		}
	}
}
