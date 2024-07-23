package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
	return c.value
}

// в задании судя по всему речь идет о методе Icrement(),
// который должен быть потокобезопасным и должен вызываться 1000 раз
// а функция concurrentSafeCounter должна проверять работу метода
func concurrentSafeCounter() int {
	wg := sync.WaitGroup{}
	counter := Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := counter.Increment()
			fmt.Println(res)
		}()
	}
	wg.Wait()
	return counter.value
}

func main() {
	fmt.Println(concurrentSafeCounter())
}
