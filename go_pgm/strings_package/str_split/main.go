package main

import (
	"fmt"
	"strings"
)

func main() {

	//Example 1: Split a String by a Space
	s := "hello world this is Go"
	result := strings.Split(s, " ")
	fmt.Println(result)

	//Example 2: Split a String by a Comma
	s1 := "apple,orange,banana,grape"
	result1 := strings.Split(s1, ",")
	fmt.Println(result1)

	//Example 3: Split by an Empty String
	s2 := "hello"
	result2 := strings.Split(s2, "")
	fmt.Println(result2)

	//Example 4: No Separator Found
	s3 := "hello"
	result3 := strings.Split(s3, ",")
	fmt.Println(result3)

	//Example 5: Splitting a String with Leading and Trailing Delimiters
	s4 := ",apple,orange,banana,"
	result4 := strings.Split(s4, ",")
	fmt.Println(result4)
}
