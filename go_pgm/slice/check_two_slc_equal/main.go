package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 6}
	if len(a) != len(b) {
		fmt.Println(false)
		return
	}
	for i := range a {
		if a[i] != b[i] {
			fmt.Println(false)
			return
		}
	}
	fmt.Println("both slices are euqal")

}
