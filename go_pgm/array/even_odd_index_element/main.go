package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // Example array
	sumEven, sumOdd := 0, 0      // Step 1: Initialize sums for even and odd indices

	// Step 2: Loop through the array
	for i, val := range arr {
		if i%2 == 0 { // Step 3: Check if the index is even
			sumEven += val // Step 4: Add the value to the sum of even indices
		} else { // Step 5: If the index is odd
			sumOdd += val // Step 6: Add the value to the sum of odd indices
		}
	}

	// Step 7: Print the results
	fmt.Printf("Sum of even indices: %d\n", sumEven) // Output the sum of even indices
	fmt.Printf("Sum of odd indices: %d\n", sumOdd)   // Output the sum of odd indices
}
