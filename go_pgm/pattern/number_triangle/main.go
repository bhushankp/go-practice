package main

import "fmt"

/*


1
1 2
1 2 3
1 2 3 4
1 2 3 4 5


*/

func main() {

	rows := 5

	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
