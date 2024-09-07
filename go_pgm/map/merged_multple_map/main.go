package main

import "fmt"

func main() {

	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"c": 3, "d": 4}
	m3 := map[string]int{"e": 5, "f": 6}

	mergedMap := map[string]int{}

	for _, m := range []map[string]int{m1, m2, m3} {
		for k, v := range m {
			mergedMap[k] = v
		}
	}

	fmt.Println(mergedMap)
}
