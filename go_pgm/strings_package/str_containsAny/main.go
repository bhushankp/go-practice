package main

import (
	"fmt"
	"strings"
)

func hasSpecialCharacterOrDigit(s string) bool {
	// Check if the string contains any digit or special character
	return strings.ContainsAny(s, "0123456789!@#$%^&*()")
}

func main() {
	// Example 1: String contains characters from the set
	s := "hello world"
	chars := "aeiou"
	result := strings.ContainsAny(s, chars)
	fmt.Println(result) // Output: true (because 'e', 'o' are in 'hello world')

	fmt.Println(hasSpecialCharacterOrDigit("hello123")) // Output: true (contains '1', '2', '3')
	fmt.Println(hasSpecialCharacterOrDigit("hello"))    // Output: false (no digit or special character)

	// Example 2: String does not contain any characters from the set
	s = "xyz"
	chars = "abc"
	result = strings.ContainsAny(s, chars)
	fmt.Println(result) // Output: false (no characters 'a', 'b', 'c' in 'xyz')

	// Example 3: Empty string or empty set
	s = ""
	chars = "abc"
	result = strings.ContainsAny(s, chars)
	fmt.Println(result) // Output: false (empty string contains no characters)

	s = "abc"
	chars = ""
	result = strings.ContainsAny(s, chars)
	fmt.Println(result) // Output: false (empty set contains no characters to find)
}
