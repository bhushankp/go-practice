package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 7}

	sum := 0

	for _, v := range slc {
		sum += v
	}
	average := float64(sum) / float64(len(slc))
	fmt.Println(average)
}
