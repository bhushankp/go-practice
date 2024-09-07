package main

import "fmt"

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	sum := 0

	for _, value := range arr {
		sum += value
	}

	fmt.Printf("%d", sum)
}
