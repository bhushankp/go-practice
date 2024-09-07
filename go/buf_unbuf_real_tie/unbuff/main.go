package main

import (
	"fmt"
	"time"
)

func playerA(c chan string) {
	for {
		// Player A throws the ball
		c <- "ball"
		fmt.Println("Player A throws the ball")

		// Wait for Player B to catch the ball
		fmt.Println("Player A is waiting for Player B to catch the ball")
		<-c
		fmt.Println("Player A got the ball back")
	}
}

func playerB(c chan string) {
	for {
		// Wait for Player A to throw the ball
		fmt.Println("Player B is waiting for Player A to throw the ball")
		<-c
		fmt.Println("Player B caught the ball")

		// Player B throws the ball back
		c <- "ball"
		fmt.Println("Player B throws the ball back")
	}
}

func main() {
	// Create an unbuffered channel
	c := make(chan string)

	// Start the players' routines
	go playerA(c)
	go playerB(c)

	// Let the game run for some time
	time.Sleep(1 * time.Second)
}
