package main

import "fmt"

func main() {

	arr := [5]int{22, 0, 0, 32, 1}

	zeroIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[zeroIndex] = arr[i]
			zeroIndex++
		}
	}

	for zeroIndex < len(arr) {
		arr[zeroIndex] = 0
		zeroIndex++
	}
	fmt.Println(arr)
}

// package main

// import (
//     "fmt"
// )

// // moveZeros moves all zeros in the array to the end while maintaining the order of non-zero elements
// func moveZeros(arr [5]int) [5]int {
//     nonZeroIndex := 0

//     // Move non-zero elements to the front of the array
//     for i := 0; i < len(arr); i++ {
//         if arr[i] != 0 {
//             arr[nonZeroIndex] = arr[i]
//             nonZeroIndex++
//         }
//     }

//     // Fill the rest of the array with zeros
//     for nonZeroIndex < len(arr) {
//         arr[nonZeroIndex] = 0
//         nonZeroIndex++
//     }

//     return arr
// }

// func main() {
//     arr := [5]int{0, 1, 0, 3, 12} // Example array
//     result := moveZeros(arr)
//     fmt.Println("Array after moving zeros:", result)
// }
