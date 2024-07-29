package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Пример структуры счетчика
type Counter struct {
	count atomic.Int64
}

// Функция для увеличения значения счетчика на 1
func (c *Counter) Increment() {
	// Ваш код для увеличения значения счетчика на 1
	c.count.Add(1)
}

// Функция для получения текущего значения счетчика
func (c *Counter) GetCount() int64 {
	// Ваш код для получения текущего значения счетчика
	return c.count.Load()
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	fmt.Println(counter.GetCount())
}
