package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 6}
	min, max := slc[0], slc[0]

	for _, v := range slc {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	fmt.Println(min, max)
}
