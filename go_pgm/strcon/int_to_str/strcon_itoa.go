package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 123

	v := strconv.Itoa(i)
	fmt.Println(v)
	fmt.Printf("type of v : %T \n", v)
	fmt.Printf("type of i : %T \n", i)

}
