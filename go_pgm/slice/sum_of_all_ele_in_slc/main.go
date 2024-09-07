package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 6}

	sum := 0

	for _, v := range slc {
		sum += v
	}

	fmt.Println("sum of all elements in slc : ", sum)
}
