package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 5}
	value := 5

	newSlc := []int{}

	for _, v := range slc {
		if v != value {
			newSlc = append(newSlc, v)
		}
	}

	fmt.Println(newSlc)
}
