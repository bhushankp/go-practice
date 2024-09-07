package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 6}

	value := 10
	contain := false

	for _, v := range slc {
		if v == value {
			contain = true
		}
	}

	if contain {
		fmt.Println("slc contain the value : ", contain)
	} else {
		fmt.Println("slc does not contain the value : ", contain)
	}
}
