package main

import (
	"fmt"
)

/*
arr := []int{1, 2, 2, 3, 4, 4, 4, 5}

	occurrences := make(map[int]int)

	for _, num := range arr {
	    occurrences[num]++
	}

	for num, count := range occurrences {
	    fmt.Printf("Element '%d' occurs %d times\n", num, count)
	}
*/
func countOccurrences(s string) map[rune]int {
	occurrence := make(map[rune]int)
	for _, char := range s {
		if char != ' ' {
			occurrence[char]++

		}
	}

	return occurrence
}

func main() {
	s := "hello   world"

	fmt.Println("Occurrences of letters:")

	for char, count := range countOccurrences(s) {
		fmt.Printf("%c=%d, ", char, count)
	}

	fmt.Println()
}
