package main

import (
	"fmt"
)

func main() {
	var amount = 100
	go func() { amount = amount + 100 }()
	fmt.Println("current amount is", amount)
}