package main

import "fmt"

var Ch = make(chan int)

func main() {

	go Sum(1, 1000)
	sum := <-Ch

	fmt.Println(sum)

}

func Sum(start, end int) int {

	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}

	Ch <- sum
	return sum

}
