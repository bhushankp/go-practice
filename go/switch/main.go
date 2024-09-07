package main

import "fmt"

func main() {
	dayOfWeek := 4

	switch dayOfWeek {
	case 1:
		fmt.Println("It's Monday")
	case 2:
		fmt.Println("It's Tuesday")
	case 3:
		fmt.Println("It's Wednesday")
	case 4:
		fmt.Println("It's Thursday")
	case 5:
		fmt.Println("It's Friday")
	case 6, 7:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Invalid day of the week")
	}
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
	//x2 := interface{}(23.2)
	// Switch with type assertion
	var x interface{} = "Hello"
	switch v := x.(type) {
	case int:
		fmt.Println("It's an int:", v)
	case string:
		fmt.Println("It's a string:", v)
	default:
		fmt.Println("Unknown type")
	}
}
