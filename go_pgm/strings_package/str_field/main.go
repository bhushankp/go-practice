package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "India is my country all indians are my brothers and sisters."
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}
}
