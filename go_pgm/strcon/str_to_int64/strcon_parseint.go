package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"

	v, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		fmt.Println("err ", err)

	} else {
		fmt.Println(v)
	}
}
