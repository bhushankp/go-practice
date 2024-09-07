package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "hello world!!"
	fmt.Println(strings.Contains(str, "hi"))
	fmt.Println(strings.Contains("hi im devil", "hi"))
}
