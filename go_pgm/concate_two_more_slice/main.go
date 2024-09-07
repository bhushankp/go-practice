package main

import "fmt"

func main() {
	a := []int{12, 34, 23, 45}
	b := []int{3, 4, 1, 5, 3}
	c := []int{55, 44, 33}

	fmt.Println("slice a : ", a)
	fmt.Println("cap slice(a) : ", cap(a))

	fmt.Println("slice b : ", b)
	fmt.Println("cap slice(b) : ", cap(b))
	fmt.Println("slice c : ", c)
	fmt.Println("cap slice(c) : ", cap(c))

	// x := []int{}
	// x = append(a, b...)
	// x = append(x, c...)

	// fmt.Printf("%v : \n", x)
	// fmt.Println("capcity of x wil be : ", cap(x))

}
