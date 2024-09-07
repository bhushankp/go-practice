package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	k := 3

	n := len(arr)
	k = k % n

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]

	}

	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i, j := k, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println(arr)

}
