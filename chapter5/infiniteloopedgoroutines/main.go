package main

import "fmt"
import "time"
import "runtime"

func main() {
	fmt.Println("Demo of infinitely looped goroutines running on GOMAXPROCS number of threads")
	threads := runtime.GOMAXPROCS(0)
	var counter int
	for i := 0; i < threads; i++ {
		go func(counter *int) {
			for {
				*counter++
			}
		}(&counter)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("value of x is", counter)
}
