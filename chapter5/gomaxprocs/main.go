package main

import "fmt"

import "runtime"

func main() {
	fmt.Println("Program has run with the value of GOMAXPROCS = ", runtime.GOMAXPROCS(0))
	newNumberOfThreads := runtime.GOMAXPROCS(0) - 1
	runtime.GOMAXPROCS(newNumberOfThreads)
	fmt.Println("Modified value of GOMAXPROCS is", runtime.GOMAXPROCS(0))
}
