package main

import "fmt"

func secondTerm(sendOnly chan<- int, value int) {
	secondPowerValue := value * 2
	fmt.Printf("Second term in geometric progression starting from  %d  is %d\n", value, secondPowerValue)
	sendOnly <- secondPowerValue
}

func thirdTerm(recieveOnly <-chan int, sendOnly chan<- int) {
	value := <-recieveOnly
	thirdPowerValue := value * 2
	sendOnly <- thirdPowerValue
}

func main() {
	initialValue := 20
	firstChannel := make(chan int, 1)
	secondChannel := make(chan int, 1)
	go secondTerm(firstChannel, initialValue)
	go thirdTerm(firstChannel, secondChannel)
	fmt.Printf("Third term in geometric progression starting from  %d  is %d\n", initialValue, <-secondChannel)
}
