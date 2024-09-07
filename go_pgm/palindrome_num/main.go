package main

import "fmt"

func main() {
	var number, original_num, reminder, reversed_num int

	fmt.Println("enter the number to check palindrome!!")
	fmt.Scanln(&number)

	original_num = number
	for number != 0 {
		reminder = number % 10
		reversed_num = reversed_num*10 + reminder
		number = number / 10

	}
	if original_num == reversed_num {
		fmt.Printf("%d is a palindrome number.\n", original_num)
	} else {
		fmt.Printf("%d is not a palindrome number.\n", original_num)
	}
}
