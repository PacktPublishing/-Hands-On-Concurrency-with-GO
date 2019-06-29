package main

import (
	"fmt"
)

func main() {
	channel := make(chan bool)
	sharedMap := make(map[string]string)
	go func() {
		sharedMap["1"] = "a" // First conflicting access.
	}()
	sharedMap["2"] = "b" // Second conflicting access.

	for k, v := range sharedMap {
		fmt.Println(k, v)
	}
	<-channel
}
