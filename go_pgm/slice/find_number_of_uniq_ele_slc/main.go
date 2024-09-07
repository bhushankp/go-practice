package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 4, 4, 3, 5, 2}

	m := map[int]bool{}

	for _, v := range slc {
		m[v] = true
	}
	fmt.Println("the number of unique elements present in slice are : ", len(m))
}
