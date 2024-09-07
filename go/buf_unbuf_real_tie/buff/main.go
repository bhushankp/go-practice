package main

import (
	"fmt"
	"time"
)

func playerA(c chan string) {
	for {
		// Player A throws the ball into the basket
		select {
		case c <- "ball":
			fmt.Println("Player A throws the ball into the basket")
		default:
			fmt.Println("Player A is waiting for space in the basket")
			time.Sleep(2 * time.Second) // Simulate waiting for space in the basket
		}
	}
}

func playerB(c chan string) {
	for {
		// Player B takes the ball from the basket
		select {
		case <-c:
			fmt.Println("Player B takes the ball from the basket")
		default:
			fmt.Println("Player B is waiting for a ball in the basket")
			time.Sleep(2 * time.Second) // Simulate waiting for a ball in the basket
		}
	}
}

func main() {
	// Create a buffered channel with a capacity of 1
	c := make(chan string, 1)

	// Start the players' routines
	go playerA(c)
	go playerB(c)

	// Let the game run for some time
	time.Sleep(5 * time.Second)
}
