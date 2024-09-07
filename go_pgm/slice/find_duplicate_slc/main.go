package main

import "fmt"

func main() {
	s := []int{1, 3, 4, 2, 2} // Step 1: Define a slice of integers `s` with a duplicate value
	m := make(map[int]bool)   // Step 2: Initialize a map `m` to track seen numbers

	// Step 3: Iterate over each element in the slice `s`
	for _, v := range s {
		// Step 4: Check if the current number `v` is already in the map `m`
		if m[v] {
			fmt.Println(v) // Step 5: If it is, print `v` as the first duplicate and exit the loop
			return
		}
		// Step 6: Otherwise, add `v` to the map `m`
		m[v] = true
	}

	// Step 7: If no duplicates are found, print `-1` (indicating no duplicates)
	fmt.Println(-1)
}
