package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "false"

	b, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("err to convert the string into bool : ", err)
	} else {
		fmt.Println(b)
	}
}
