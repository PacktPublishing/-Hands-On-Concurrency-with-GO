package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo of cancellation of goroutine")
	ctx, cancel := context.WithCancel(context.Background())
	commChannel := make(chan struct{}, 1)
	go longRunningWork(ctx, commChannel)
	commChannel <- struct{}{}
	time.Sleep(5 * time.Second)
	defer cancel()
}

func longRunningWork(ctx context.Context, reciever <-chan struct{}) {
	fmt.Println("going to start doing some long running works")
	for true {
		select {
		case <-ctx.Done():
			fmt.Println("I am being asked to stop")
			return
		case <-reciever:
			fmt.Println("Doing some long running work")
		default:
		}
	}
	fmt.Println("Done the long running work")
}
