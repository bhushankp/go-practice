package main

import (
	"fmt"
)

func main() {

	var first, second int //5 6

	fmt.Println("enter the first number")
	fmt.Scanln(&first)
	fmt.Println("enter the second number")
	fmt.Scanln(&second)
	fmt.Println("before changing the first number", first)
	fmt.Println("before changing the second number", second)

	first = first - second
	second = first + second
	first = second - first

	fmt.Println("after changing the first number", first)
	fmt.Println("after changing the second number", second)

}
