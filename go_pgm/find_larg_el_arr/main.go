package main

import "fmt"

func main() {

	arr := [5]int{12, 44, 22, 11, 4}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
