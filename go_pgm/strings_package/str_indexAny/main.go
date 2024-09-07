package main

import (
	"fmt"
	"strings"
)

func main() {

	//Example 1: Finding Any Character from a Set
	s := "hello, world!"
	index := strings.IndexAny(s, "aeiou")
	fmt.Println(index)

	//Example 2: Characters Not Found
	s1 := "hello, world!"
	index1 := strings.IndexAny(s1, "xyz")
	fmt.Println(index1)

	//Example 3: Multiple Characters

	s2 := "golang"
	index2 := strings.IndexAny(s2, "aeiou")
	fmt.Println(index2)

	//Example 4: Searching for a Single Character
	s3 := "golang"
	index3 := strings.IndexAny(s3, "n")
	fmt.Println(index3)

	//
}
