package main

import "fmt"

func main() {

	fmt.Println("enter the three numbers!!")
	var a int
	fmt.Scanln(&a)
	var b int
	fmt.Scanln(&b)
	var c int
	fmt.Scanln(&c)

	if a > b && a > c {
		fmt.Println("first number is the largest number!!")
	} else if b > a && b > c {
		fmt.Println("second number is the largest number!!")
	} else {
		fmt.Println("third number is the largest number!!")
	}

}
