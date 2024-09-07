package main

import "fmt"

func main() {

	arr := [5]int{3, 2, 1, 4, 5}

	// for i := 0; i < len(arr); i++ {
	// 	arr2[i] += arr[len(arr)-1-i]
	// }
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("Sorted array", arr)
}
