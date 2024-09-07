package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 4, 5} // Example array with the number 3 missing
	n := 5                    // The maximum number in the complete sequence

	total := n * (n + 1) / 2 // Step 1: Calculate the expected total sum of numbers from 1 to n
	sum := 0                 // Step 2: Initialize a variable to hold the sum of the numbers in the array

	// Step 3: Loop through the array to calculate the sum of its elements
	for i := 0; i < len(arr); i++ {
		sum += arr[i] // Step 4: Add each element of the array to sum
	}

	// Step 5: Output the difference between the expected total and the actual sum
	missingNumber := total - sum
	fmt.Println("The missing number is:", missingNumber) // Output the result
}
