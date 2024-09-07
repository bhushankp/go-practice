package main

import "fmt"

func main() {

	slc := []int{1, 2, 3, 2, 1}

	isPalindrome := true

	for i, j := 0, len(slc)-1; i < j; i, j = i+1, j-1 {
		if slc[i] != slc[j] {
			isPalindrome = false
			break
		}

	}
	if isPalindrome {
		fmt.Println("this is palindrome slice")
	} else {
		fmt.Println("this is not palindrome slice")
	}
}
