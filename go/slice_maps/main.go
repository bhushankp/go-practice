package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5}

	mp := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println("****Slice****")

	for i, v := range slc {
		fmt.Printf("Index : %d	, Value : %d \n", i, v)
	}
	fmt.Println("****Map****")
	for k, v := range mp {
		fmt.Printf("Key : %s	, Value : %d \n", k, v)
	}

}
