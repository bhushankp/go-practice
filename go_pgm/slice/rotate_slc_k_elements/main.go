package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} // Initial slice
	k := 2                    // Number of positions to rotate
	n := len(s)               // Length of the slice

	k = k % n // Normalize k to be within the length of the array

	// Reverse the first part of the slice (from start to n-k)
	for i, j := 0, n-k-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	// Reverse the second part of the slice (from n-k to end)
	for i, j := n-k, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	// Reverse the entire slice
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s) // Output: [4 5 1 2 3]
}
