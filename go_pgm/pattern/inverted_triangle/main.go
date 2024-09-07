package main

import "fmt"

/*

* * * * *
* * * *
* * *
* *
*


 */

func main() {

	n := 5
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
