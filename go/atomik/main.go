package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

var requestCount int64 // Global counter for the number of requests

// handler function to respond to HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Increment the request count atomically
	atomic.AddInt64(&requestCount, 1)

	// Respond with the current request count
	fmt.Fprintf(w, "Request Count: %d\n", atomic.LoadInt64(&requestCount))
}

func main() {
	http.HandleFunc("/", handler) // Register the handler function for the root URL

	// Start the web server
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// // The atomic package in Go provides low-level, thread-safe operations for basic types like integers, booleans, and pointers.
// // It allows you to perform atomic operations such as adding, loading, storing, and swapping values without the need for locks.
// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func main() {
// 	// Declare an int64 variable for atomic operations
// 	var counter int64

// 	// Atomic Add: Increment the counter by 10
// 	atomic.AddInt64(&counter, 10)
// 	fmt.Println("After Add:", counter) // Should print: 10

// 	// Atomic Load: Safely read the value of the counter
// 	currentValue := atomic.LoadInt64(&counter)
// 	fmt.Println("Current Value:", currentValue) // Should print: 10

// 	// Atomic Store: Set the counter to a new value
// 	atomic.StoreInt64(&counter, 25)
// 	fmt.Println("After Store:", counter) // Should print: 25

// 	// Atomic Swap: Swap the current value with a new value
// 	swappedValue := atomic.SwapInt64(&counter, 100)
// 	fmt.Println("Swapped Value:", swappedValue)       // Should print: 25 (old value)
// 	fmt.Println("Current Value After Swap:", counter) // Should print: 100 (new value)
// }

// // AddInt64: Atomically adds a value to counter.
// // LoadInt64: Atomically loads the current value of counter.
// // StoreInt64: Atomically sets counter to a new value.
// // SwapInt64: Atomically swaps the current value with a new one and returns the old value.
