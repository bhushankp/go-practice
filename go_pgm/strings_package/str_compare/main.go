package main

import (
	"fmt"
	"strings"
)

func main() {

	result := strings.Compare("a", "a")
	result1 := strings.Compare("a", "b")
	result2 := strings.Compare("b", "a")

	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)

	if result == 0 || result1 == 0 || result2 == 0 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are Not equal")

	}
}
