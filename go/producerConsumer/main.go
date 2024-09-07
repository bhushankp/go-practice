package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5
const numProducers = 2
const numConsumers = 3
const numItems = 10

var buffer = make(chan int, bufferSize)
var wg sync.WaitGroup

func producer(id int) {
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		item := i + id*numItems
		buffer <- item
		fmt.Printf("Producer %d produced item %d\n", id, item)
		time.Sleep(time.Millisecond * 500) // Simulate some work
	}
	close(buffer)
}

func consumer(id int) {
	defer wg.Done()
	for i := 0; i < numItems; i++ {
		item := <-buffer
		fmt.Printf("Consumer %d consumed item %d\n", id, item)
		time.Sleep(time.Millisecond * 300) // Simulate some work
	}
}

func main() {
	// Add producers and consumers to the wait group
	wg.Add(numProducers + numConsumers)

	// Start producers
	for i := 0; i < numProducers; i++ {
		go producer(i)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		go consumer(i)
	}

	// Wait for all producers and consumers to finish
	wg.Wait()

	// Close the buffer channel to signal that no more items will be produced
	close(buffer)
}
