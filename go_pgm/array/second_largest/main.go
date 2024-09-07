package main

import "fmt"

func main() {
	arr := [5]int{4, 4, 4, 4, 4}

	first, second := arr[0], -1

	for i := 1; i < len(arr); i++ {
		if arr[i] > first {
			second = first
			first = arr[i]
		} else if arr[i] > second && arr[i] < first {
			second = arr[i]
		}
	}
	if second == first {
		fmt.Println("No second max")
	}
	if second != -1 {
		fmt.Println(second)
	} else {
		fmt.Println("all elements are same no second found ")
	}
}

// package main

// import "fmt"

// func secondLarg(arr [5]int) int {

// 	first, second := arr[0], -1
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > first {
// 			second = first
// 			first = arr[i]
// 		} else if arr[i] > second && arr[i] < first {
// 			second = arr[i]
// 		}
// 	}
// 	if second == first {
// 		fmt.Println("no second large element")
// 		return -1
// 	}
// 	return second
// }

// func main() {

// 	arr := [5]int{28, 56, 3, 4, 48}

// 	if secondLarg(arr) != -1 {

// 		fmt.Println(secondLarg(arr))
// 	}

// }
