package main

import "fmt"

func main() {
	var num [100]int

	var temp, sum int
	var avg float32

	fmt.Print("enter the number of elements!!")
	fmt.Scanln(&temp)
	for i := 0; i < temp; i++ {
		fmt.Print("enter the numbers!")
		fmt.Scanln(&num[i])
		sum += num[i]
	}
	avg = (float32)(sum / temp)
	fmt.Printf("avg of the above elements are : %0.1f ", avg)
}
