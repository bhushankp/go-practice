package main

import "fmt"

func main() {

	slc := []int{1, 2, 2, 3, 4, 5, 6, 2, 1}
	m := map[int]bool{}
	uniq := []int{}

	for _, v := range slc {
		if !m[v] {
			m[v] = true
			uniq = append(uniq, v)
		}
	}
	fmt.Println(uniq)
}
