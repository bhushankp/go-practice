package main

import "fmt"

func main() {
	slc := []int{1, 2, 3, 4, 5, 5, 2, 5, 6}

	countOccur := map[int]int{}

	for _, v := range slc {
		countOccur[v]++
	}
	fmt.Println(countOccur)

	for key, val := range countOccur {
		fmt.Printf("%d is ouccr %d time\n", key, val)
	}

}
