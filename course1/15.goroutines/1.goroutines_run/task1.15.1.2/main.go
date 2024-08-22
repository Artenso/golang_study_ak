package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем новый тикер с интервалом 1 секунда
	ticker := time.NewTicker(1 * time.Second)
	data := NotifyEvery(ticker, 5*time.Second, "Таймер сработал")
	for v := range data {
		fmt.Println(v)
	}
	fmt.Println("Программа завершена")
}

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	res := make(chan string)
	timer := time.NewTimer(d)

	go func() {
		defer close(res)
		for {
			select {
			case <-ticker.C:
				res <- message
			case <-timer.C:
				ticker.Stop()
				return
			}
		}
	}()

	return res
}
