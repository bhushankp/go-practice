package main

import (
	"fmt"
)

// divide function that panics if division by zero occurs
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero") // Trigger a panic
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r) // Handle the panic
		}
		fmt.Println("can not devide by zero!!")
	}()

	// Attempt to divide by zero
	result := divide(10, 0)        // This will cause a panic
	fmt.Println("Result:", result) // This line will not be executed
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan int) // Create a channel of type int

// 	go func() {
// 		// Simulate some work
// 		time.Sleep(1 * time.Second)
// 		close(ch) // Close the channel after work is done
// 	}()

// 	// Wait for the channel to be closed
// 	time.Sleep(2 * time.Second)

// 	// Attempt to send a value on the closed channel
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from panic:", r) // Handle the panic
// 		}
// 	}()

// 	ch <- 42 // This will cause a panic

// 	fmt.Println("This line will not be executed") // This line will NOT be executed
// }
