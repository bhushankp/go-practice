package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	checkDup := map[int]bool{}
	for _, val := range arr {
		if checkDup[val] {
			fmt.Println("array contains duplicate")
			return
		}
		checkDup[val] = true
	}
	fmt.Println("array does not contain dupliates")
}
