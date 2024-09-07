package main

import "fmt"

func main() {
	fmt.Println("enter the number!!")
	var a int
	fmt.Scanln(&a)

	for i := 1; i <= 10; i++ {
		fmt.Println(a, " X ", i, " = ", a*i)

	}
}
