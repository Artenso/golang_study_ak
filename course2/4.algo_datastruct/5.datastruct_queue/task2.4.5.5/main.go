package main

import "fmt"

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type ringBuffer struct {
	buff chan int
}

func NewRingBuffer(size int) CircuitRinger {
	return &ringBuffer{
		buff: make(chan int, size),
	}
}

func (r *ringBuffer) Add(val int) {
	select {
	case r.buff <- val:
		return
	default:
		<-r.buff
		r.buff <- val
	}
}

func (r *ringBuffer) Get() (int, bool) {
	select {
	case val := <-r.buff:
		return val, true
	default:
		return 0, false
	}

}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4) // Перезаписывает значение 1
	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val) // Выводит: 2, 3, 4
	}
	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty") // Выводит: Buffer is empty
	}
}
