package main

import "fmt"

func main() {
	channel := make(chan bool)
	sharedMap := make(map[string]string)
	go func() {
		sharedMap["1"] = "a" // First conflicting access.
		channel <- true
	}()
	sharedMap["2"] = "b" // Second conflicting access.
	<-channel
	for k, v := range sharedMap {
		fmt.Println(k, v)
	}
}
