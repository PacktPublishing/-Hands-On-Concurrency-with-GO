package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("First statement of main")
	greet("Foo")
	go greet("Foo in Goroutine")
	time.Sleep(1 * time.Second)
}

func greet(greetee string) {
	fmt.Printf("Hello %s\n", greetee)
}
