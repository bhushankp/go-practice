package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "1236"

	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("error to convert string into int", err)
	} else {
		fmt.Println(i)
	}

}
