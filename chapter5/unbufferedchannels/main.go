package main

import (
	"fmt"
	"time"
)

func main() {

	stringPipeline := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		stringPipeline <- "ping"
	}()
	fmt.Println("Waiting for the message")
	msg := <-stringPipeline
	fmt.Println(msg)
}
