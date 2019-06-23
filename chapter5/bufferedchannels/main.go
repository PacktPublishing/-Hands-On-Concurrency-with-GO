package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "first"
	messages <- "second"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
