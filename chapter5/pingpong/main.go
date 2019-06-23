package main

import (
	"fmt"
	"time"
)

func main() {
	turn := make(chan struct{})
	var ball struct{}
	go player("pong ", turn)
	go player("ping", turn)

	turn <- ball // Game on; Toss the ball
	time.Sleep(2 * time.Second)
	<-turn // game over; grab the ball

}

func player(action string, turn chan struct{}) {
	for {
		thisTurn := <-turn
		fmt.Println(action)
		time.Sleep(100 * time.Millisecond)
		turn <- thisTurn
	}

}
