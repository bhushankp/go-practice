package main

import "fmt"

func main() {

	a, b := 10, 20
	op := "-"
	result := 0

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b

	default:
		result = 0
	}
	fmt.Println(result)
}
