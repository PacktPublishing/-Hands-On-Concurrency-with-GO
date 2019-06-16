package main

import "fmt"
import "time"
import "runtime"

func main() {
	fmt.Println("Demo of infinitely looped goroutines which are self yielded")
	threads := runtime.GOMAXPROCS(0)
	var counter int
	for i := 0; i < threads; i++ {
		go func(counter *int) {
			for {
				*counter++
				runtime.Gosched()
			}
		}(&counter)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("value of x is", counter)
}
