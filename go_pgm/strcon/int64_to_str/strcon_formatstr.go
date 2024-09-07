package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := int64(123)
	str := strconv.FormatInt(i, 10)
	fmt.Println(str)
}
