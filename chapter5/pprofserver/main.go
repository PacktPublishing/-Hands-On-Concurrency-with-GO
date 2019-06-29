package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	turn := make(chan struct{})
	var ball struct{}
	go player("pong ", turn)
	go player("ping", turn)

	turn <- ball // Game on; Toss the ball
	fmt.Scanln()

}

func player(action string, turn chan struct{}) {
	for {
		thisTurn := <-turn
		fmt.Println(action)
		time.Sleep(100 * time.Millisecond)
		turn <- thisTurn
	}

}
