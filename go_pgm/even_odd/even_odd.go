package main

import "fmt"

func main() {

	var number int

	fmt.Println("enter the that you check even or odd")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("the given number is even")
	} else {
		fmt.Println("the given numver is odd")
	}
}
