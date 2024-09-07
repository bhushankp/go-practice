package main

import "fmt"

func main() {

	arr := [6]int{1, 1, 2, 2, 3, 3}
	if len(arr) == 0 {
		fmt.Println("no duplicate found ")
		return
	}

	uniqueLength := 1

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[uniqueLength] = arr[i]
			uniqueLength++
		}
	}

	fmt.Println("number of dupblicate element", uniqueLength)
	fmt.Println("array after removing duplicates", arr[:uniqueLength])

}
