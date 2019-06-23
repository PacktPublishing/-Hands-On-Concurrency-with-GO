package main

import "fmt"
import "time"

func doWork(done chan bool) {
	fmt.Println("Doing work...")
	time.Sleep(1 * time.Second)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go doWork(done)
	<-done
}
