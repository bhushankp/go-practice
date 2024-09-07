package main

import "fmt"

func main() {

	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	sort := isSorted(arr)
	if sort {
		fmt.Println("The array is sorted")
	} else {
		fmt.Println("array is not sorted")
	}
}

func isSorted(arr [5]int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
