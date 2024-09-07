package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 2, 6, 7, 3}

	m := map[int]bool{}

	uniqSlc := []int{}

	for _, item := range slc {
		if _, exists := m[item]; exists {
			m[item] = true
			uniqSlc = append(uniqSlc, item)
		}
	}
	fmt.Println(uniqSlc)
}
