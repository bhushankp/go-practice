package main

import "fmt"

func main() {

	arr := [5]rune{'a', 'b', 'c', 'e', 'i'}
	vovels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	countVovels := 0

	for _, char := range arr {
		if vovels[char] {
			countVovels++
		}
	}
	fmt.Println("number of vovels present in this array : ", countVovels)
}
