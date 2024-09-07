package main

import (
	"fmt"
	"strconv"
)

func main() {

	b := true

	str := strconv.FormatBool(b)
	fmt.Println(str)
}
