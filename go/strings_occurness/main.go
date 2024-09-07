package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Go is fun. Go is powerful. Go is concise."
	substring := "Go"

	count := strings.Count(str, substring)
	fmt.Printf("The substring '%s' occurs %d times.\n", substring, count)
}
