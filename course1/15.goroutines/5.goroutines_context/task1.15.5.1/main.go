package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
	/* Output:
	   context deadline exceeded
	   time after exceeded
	*/
}

func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {
	var cancel context.CancelFunc
	ctx1, cancel := context.WithDeadline(ctx, time.Now().Add(contextDeadline))
	defer cancel()

	select {
	case <-ctx1.Done():
		return "context deadline exceeded"
	case <-time.After(timeAfter):
		return "time after exceeded"

	}
}
