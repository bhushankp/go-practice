package main

import "fmt"

/*

*********
 *******
  *****
   ***
    *



*/

func main() {
	var rows int
	fmt.Print("enter the rows of reverse paramid : ")
	fmt.Scan(&rows)

	for i := rows; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
