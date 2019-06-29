package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo of cancellation of goroutine")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	trigger := make(chan struct{}, 1)
	callback := make(chan struct{}, 1)
	status := "READY"
	go longRunningWork(ctx, trigger, callback, &status)
	trigger <- struct{}{}
	time.Sleep(2 * time.Second)
	cancel()
	<-callback
	fmt.Printf("Status Value is %s\n", status)
}

func longRunningWork(ctx context.Context, trigger <-chan struct{}, callback chan<- struct{}, status *string) {
	fmt.Println("Going to start doing some long running works")
	for true {
		select {
		case <-ctx.Done():
			fmt.Println("I am being asked to stop")
			*status = "TERMINATED"
			callback <- struct{}{}
			return
		case <-trigger:
			fmt.Printf("Status value is %s\n", *status)
			fmt.Println("Doing some long running work")
			*status = "EXECUTING"
			fmt.Printf("Status value is %s\n", *status)

		default:
		}
	}
	fmt.Println("Done the long running work")
}
