package main

import "fmt"

func main() {

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 3}

	compare := true

	if len(m1) != len(m2) {
		compare = false
	}

	for k, v := range m1 {
		if v2, exists := m2[k]; !exists || v != v2 {
			compare = false
		}
	}

	if compare {
		fmt.Println("both maps are equal")
	} else {
		fmt.Println("both maps not are equal")
	}
}
