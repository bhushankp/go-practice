package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 6, 7}

	// var evenSlc []int
	evenSlc := []int{}
	for _, v := range slc {
		if v%2 == 0 {
			evenSlc = append(evenSlc, v)
		}
	}
	fmt.Println(evenSlc)
}
