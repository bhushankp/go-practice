package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 7}

	n := 7

	total := n * (n + 1) / 2

	sum := 0

	for _, v := range slc {
		sum += v
	}
	result := total - sum
	fmt.Println(result)
}
