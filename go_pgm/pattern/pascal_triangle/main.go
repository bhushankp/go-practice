package main

import (
	"fmt"
)

/*

1
1 1
1 2 1
1 3 3 1
1 4 6 4 1


*/

func main() {
	var n int
	fmt.Print("enter the rows of pascal triangle : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		number := 1
		for j := 1; j <= i; j++ {
			fmt.Print(number)
			number = number * (i - j) / j
		}
		fmt.Println()
	}

}

// func main() {
// 	n := 5 // Number of rows in Pascal's Triangle
// 	for i := 0; i < n; i++ {
// 		number := 1 // The first element in any row is always 1
// 		for j := 0; j <= i; j++ {
// 			fmt.Printf("%d ", number)
// 			// Calculate the next element in the row
// 			number = number * (i - j) / (j + 1)
// 		}
// 		fmt.Println()
// 	}
// }
