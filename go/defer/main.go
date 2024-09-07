package main

import "fmt"

func main() {

	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 3")
	defer fmt.Println("Deferred statement 2")

	// Other statements in the function
	fmt.Println("Regular statement 1")
	fmt.Println("Regular statement 2")

}
