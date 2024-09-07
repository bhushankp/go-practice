package main

import "fmt"

func main() {

	// var a [3]int = [3]int{5, 2, 3}
	// l := len(a)
	// var b [3]int
	// for i := l - 1; i >= 0; i-- {
	// 	b[l-i-1] += a[i]
	// }
	// fmt.Println(b)

	arr := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(arr); i++ {
	// 	arr2[i] += arr[len(arr)-1-i]
	// }
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Printf("%d", arr)
}
