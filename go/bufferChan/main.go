package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 3)

	var wg sync.WaitGroup

	wg.Add(2)

	//sender goroutine
	go func() {
		defer wg.Done()
		value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, v := range value {
			fmt.Println("sending value from channel", v)
			ch <- v //Does not block until the buffer is full\
		}
		close(ch)
	}()

	//receiver goroutine
	go func() {
		defer wg.Done()
		for {
			receivedValue, ok := <-ch
			if !ok {
				fmt.Println("channel is clossed, no more data")
				return
			}
			fmt.Println("received Value from channel", receivedValue)

		}

	}()

	wg.Wait()
	fmt.Println("wating for goroutine finish")
}
