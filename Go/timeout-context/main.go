package main

import (
	"context"
	"time"
)

func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case <-ctx.Done():
		println("timeout")
	case v := <-ch:
		println("valor:", v)
	}
}

func main() {
	withTimeout()
}
