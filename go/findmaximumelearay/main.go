package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	fmt.Println("max element of array", max)
}
