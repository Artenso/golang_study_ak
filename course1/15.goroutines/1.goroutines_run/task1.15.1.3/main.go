package main

import (
	"fmt"
	"sync"
	"time"
)

// Job is a job for worker
type Job struct {
	id int
}

// WorkerPool is pool of workers
type WorkerPool struct {
	jobs    chan Job
	results chan int
}

// NewWorkerPool creates new workers pool
func NewWorkerPool(numWorkers int, queue chan Job) *WorkerPool {
	return &WorkerPool{
		jobs:    queue,
		results: make(chan int, len(queue)),
	}
}

// Run starts workers
func (wp *WorkerPool) Run(numWorkers int) {
	wg := sync.WaitGroup{}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for job := range wp.jobs {
				// some work here
				res := job.id * 2
				time.Sleep(time.Second)
				// send results to chan
				wp.results <- res
			}
		}(i)
	}

	go func() {
		defer close(wp.results)
		wg.Wait()
	}()
}

// AddJob adds new job to queue
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

func main() {
	numWorkers := 3
	queue := make(chan Job, 3)

	workerPool := NewWorkerPool(numWorkers, queue)

	workerPool.Run(numWorkers)

	go func() {
		defer close(queue)
		for i := 1; i <= 15; i++ {
			job := Job{id: i}
			workerPool.AddJob(job)
		}
	}()

	for res := range workerPool.results {
		fmt.Printf("Result: %d\n", res)
	}
}
