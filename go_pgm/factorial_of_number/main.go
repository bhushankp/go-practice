package main

import (
	"fmt"
)

// Iterative function to compute factorial
func factorialIterative(n int) int {
	if n < 0 {
		return 0 // Handle negative numbers
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func main() {
	num := 5 // Example number
	result := factorialIterative(num)
	fmt.Printf("Factorial of %d is %d\n", num, result)
}
