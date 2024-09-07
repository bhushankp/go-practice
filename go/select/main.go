package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	channel1 := make(chan string)
	channel2 := make(chan string)

	// Goroutine 1 sending a message to channel1 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "Message from Goroutine 1"
	}()

	// Goroutine 2 sending a message to channel2 after 4 seconds
	go func() {
		time.Sleep(3 * time.Second)
		channel2 <- "Message from Goroutine 2"
	}()

	// Use select to wait for the first channel that becomes ready
	select {
	case message1 := <-channel1:
		fmt.Println(message1)
	case message2 := <-channel2:
		fmt.Println(message2)
	}

	fmt.Println("Program completed.")
}
