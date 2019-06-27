package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo of cancellation of goroutine")
	ctx, cancel := context.WithCancel(context.Background())

	go longRunningWork(ctx)
	time.Sleep(2 * time.Second)

	cancel()
	fmt.Scanln()
}

func longRunningWork(ctx context.Context) {
	fmt.Println("going to start doing some long running works")
	for true {
		select {
		case <-ctx.Done():
			fmt.Println("I am being asked to stop")
			return
		default:
			fmt.Println("doing a long running work")
		}
	}
	fmt.Println("Done the long running work")
}
