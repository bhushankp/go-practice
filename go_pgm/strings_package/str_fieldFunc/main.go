package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//Example 1: Splitting on Whitespace Characters
	s := "Hello, \tworld!\nThis is Go."
	result := strings.FieldsFunc(s, unicode.IsSpace)
	fmt.Println(result)

	//Example 2: Splitting on Specific Characters
	s1 := "go1|python2|java3"
	result1 := strings.FieldsFunc(s1, func(r rune) bool {
		return r == '|' || r == '1' || r == '2' || r == '3'
	})
	fmt.Println(result1)

	//Example 3: Splitting on Vowels
	s2 := "separate on vowels"
	result2 := strings.FieldsFunc(s2, func(r rune) bool {
		return strings.ContainsRune("aeiouAEIOU", r)
	})
	fmt.Println(result2)
}
