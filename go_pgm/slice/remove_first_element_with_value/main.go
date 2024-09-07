package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 5, 3}

	value := 3

	newSlc := []int{}
	found := false

	for _, v := range slc {
		if v == value && !found {
			found = true
			continue
		}
		newSlc = append(newSlc, v)
	}
	fmt.Println(newSlc)
}
