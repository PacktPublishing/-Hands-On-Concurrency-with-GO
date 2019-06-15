package main

import (
	"fmt"
)

func main() {
	fmt.Println("Table of Multiples for 2 till 10 are")

	go func(value int) {
		for index := 1; index <= 10; index++ {
			fmt.Printf("%d times %d is %d\n", index, value, value*index)
		}
	}(2)
	fmt.Scanln()
	fmt.Println("Done!")
}
