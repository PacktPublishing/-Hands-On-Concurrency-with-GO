package main

import "time"
import "fmt"

func main() {

	firstChannel := make(chan string)
	secondChannel := make(chan string)
	go func() {

		time.Sleep(1 * time.Second)
		firstChannel <- "one"

	}()
	go func() {

		time.Sleep(1 * time.Second)
		secondChannel <- "two"

	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-firstChannel:
			fmt.Println("I am", msg1)
		case msg2 := <-secondChannel:
			fmt.Println("I am", msg2)
		}
	}

}
