package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // unbuffered channel

	// Sender goroutine
	wg.Add(2)
	go func() {
		defer wg.Done()
		value := 42
		fmt.Println("Sending value to the channel:", value)
		ch <- value // Blocks until a receiver is ready
		//fmt.Println("Sent value to the channel")
	}()

	// Receiver goroutine
	go func() {
		defer wg.Done()
		receivedValue := <-ch // Blocks until a sender is ready
		fmt.Println("Received value from the channel:", receivedValue)
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	//close(ch)
}
