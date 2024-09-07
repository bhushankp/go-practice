package main

import "fmt"

/*

        *
      * * *
    * * * * *
  * * * * * * *
* * * * * * * * *
  * * * * * * *
    * * * * *
      * * *
        *


*/

func main() {
	rows := 5

	for i := 1; i <= rows; i++ {
		for j := i; j < rows; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for j := rows - 1; j >= 1; j-- {
		for k := rows; k > j; k-- {
			fmt.Print(" ")
		}
		for l := 1; l <= 2*j-1; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
