package main

import "fmt"

type Calculator struct{}

//type BigCalculator struct{}
//same name two function not allowed thats why methods instace created
func Add(x, y int) int {
	return x + y
}

// func (c BigCalculator) Add(x, y int) int {
// 	return x + y
// }
func (c Calculator) Add(x, y int) int {
	return x + y
}

func main() {

	//function are standalone
	//function are called independetly
	result := Add(5, 10)
	fmt.Println(result)

	//methods are associated with specific type
	//methdos are called instance of type
	c := Calculator{}
	resultt := c.Add(2, 4)
	fmt.Println(resultt)

}
