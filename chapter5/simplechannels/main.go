package main

import (
	"fmt"
)

func main() {

	stringPipeline := make(chan string)

	go func() {
		stringPipeline <- "ping"
	}()

	msg := <-stringPipeline
	fmt.Println(msg)
}
