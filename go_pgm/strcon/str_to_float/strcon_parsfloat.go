package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "234.3"
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println(" error to convert the string into float : ", err)
	} else {
		fmt.Println(f)
	}

}
