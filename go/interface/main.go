package main

import "fmt"

// 1. Definition of an Interface:
// An interface is a type that specifies a set of method signatures.
type Speaker interface {
	Speak() string
}

// 2. Implementation of an Interface:
// Any type that implements the methods defined by an interface is said to "implement" that interface implicitly.
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I am " + d.Name
}

// 3. Polymorphism:
// Interfaces enable polymorphism, allowing different types to be used interchangeably.
func Greet(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	// 4. Using the interface with different types:
	var p Speaker = Person{Name: "John"}
	var d Speaker = Dog{Name: "Buddy"}

	// Both Person and Dog satisfy the Speaker interface and can be used interchangeably.
	Greet(p) // Output: Hello, my name is John
	Greet(d) // Output: Woof! I am Buddy

	// 5. Empty Interface:
	// The empty interface can hold values of any type.
	var i interface{}
	i = "A string value"
	fmt.Println(i) // Output: A string value

	i = 42
	fmt.Println(i) // Output: 42

	// 6. Type Assertion:
	// Type assertion is used to extract the concrete value from an interface variable.
	if str, ok := i.(int); ok {
		fmt.Println("i holds an integer:", str)
	}

	// 7. Abstraction:
	// Interfaces allow you to abstract away the implementation details, focusing on what a type can do.
	// In this example, the Greet function works with any type that satisfies the Speaker interface,
	// regardless of whether it's a Person, Dog, or any other type that implements Speak().
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geomentry interface {
// 	area() float64
// 	paramiter() float64
// }

// func (r Rectangle) area() float64 {

// 	return r.height * r.weight

// }

// func (r Rectangle) paramiter() float64 {
// 	return 2*r.height + 2*r.weight
// }

// func (c Circle) area() float64 {

// 	return math.Pi * (float64)(c.redius*c.redius)
// }

// func (c Circle) paramiter() float64 {
// 	return 2 * math.Pi * float64(c.redius)
// }

// type Circle struct {
// 	redius int
// }

// type Rectangle struct {
// 	height, weight float64
// }

// func main() {

// 	var g geomentry = Rectangle{height: 10, weight: 5}
// 	fmt.Println(g.area())      // 50
// 	fmt.Println(g.paramiter()) // 30

// 	g = Circle{redius: 7}
// 	fmt.Println(g.area())      // 15π
// 	fmt.Println(g.paramiter()) // 16π
// 	fmt.Printf("Type %T\n", g) // Type main.Circle

// 	// var rec Rectangle = Rectangle{5, 10}
// 	// var cir Circle = Circle{3}

// 	// 	print("The Area of rectangle is: ", rec.area())
// 	// 	println("The Parameter of rectangle is:", rec.paramiter())

// 	// 	print("The Area of circle is: ", cir.area())
// 	// 	println("The Parameter of circle is:", cir.paramiter())
// 	// }
// }
