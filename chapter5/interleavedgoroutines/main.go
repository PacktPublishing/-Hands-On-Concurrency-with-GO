package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ping Pong Demo")
	go pinger()
	go ponger()
	fmt.Scanln()
	fmt.Println("done")
}

func pinger() {
	for index := 0; index < 10; index++ {
		fmt.Println("ping")
		time.Sleep(100 * time.Millisecond)
	}

}

func ponger() {
	for index := 0; index < 10; index++ {
		fmt.Println("pong")
		time.Sleep(100 * time.Millisecond)
	}

}
