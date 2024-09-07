package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 2, 1} // Example array that is a palindrome
	isPalindrome := true         // Assume it is a palindrome initially

	// Loop through half the array
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] { // Compare front and back
			isPalindrome = false // Set to false if they don't match
			break                // Exit the loop early
		}
	}

	// Print the result
	if isPalindrome {
		fmt.Println("The array is a palindrome.")
	} else {
		fmt.Println("The array is not a palindrome.")
	}
}
