package main

import (
	"fmt"
	"strings"
)

func main() {
	//Example 1: Replace All Occurrences wiht -1
	s := "hello world, hello universe"
	result := strings.Replace(s, "hello", "hi", -1)
	fmt.Println(result)

	//Example 2: Replace Only the First Occurrence
	s1 := "hello world, hello universe"
	result1 := strings.Replace(s1, "hello", "hi", 1)
	fmt.Println(result1)

	//Example 3: Replace a Specific Number of Occurrences
	s2 := "one two three two four two"
	result2 := strings.Replace(s2, "two", "2", 2)
	fmt.Println(result2)

	//Example 4: No Occurrences Found
	s3 := "hello world"
	result3 := strings.Replace(s3, "goodbye", "hi", -1)
	fmt.Println(result3)
}
