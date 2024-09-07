package main

import (
	"fmt"
	"strconv"
)

func main() {

	flt := 234.4

	str := strconv.FormatFloat(flt, 'f', 3, 32)
	str1 := strconv.FormatFloat(flt, 'e', -1, 32)
	str2 := strconv.FormatFloat(flt, 'g', 3, 32)
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)

}

// Convert using different formats
// s1 := strconv.FormatFloat(f, 'f', 2, 64)  // Decimal, 2 digits after decimal point
// s2 := strconv.FormatFloat(f, 'e', -1, 64) // Exponent notation, automatic precision
// s3 := strconv.FormatFloat(f, 'g', 4, 64)  // Shortest representation with 4 significant digits

// fmt.Println("Decimal format (2 digits):", s1) // Output: 123.46
// fmt.Println("Exponent format:", s2)           // Output: 1.23456789e+02
// fmt.Println("Shortest representation (4 digits):", s3) // Output: 123.5
