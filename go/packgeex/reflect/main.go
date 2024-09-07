package main

import (
	"fmt"
	"reflect"
)

func main() {
	variable := interface{}(32)

	//type of value
	typeOfVariable := reflect.TypeOf(variable)
	fmt.Println("type of value : ", typeOfVariable)
	//value of value
	valueOfVariable := reflect.ValueOf(variable)
	fmt.Println("type of value : ", valueOfVariable)

}

// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	type Person struct {
// 		Name string
// 		age  int64
// 	}

// 	p := Person{Name: "John", age: 30}

// 	typeOfPerson := reflect.TypeOf(p)
// 	valueOfPerson := reflect.ValueOf(p)

// 	fmt.Println(typeOfPerson)                                // prints: main.Person
// 	fmt.Printf("Kind of type is %s\n", valueOfPerson.Kind()) // prints: Struct

// 	for i := 0; i < valueOfPerson.NumField(); i++ {
// 		field := valueOfPerson.Field(i)
// 		fieldName := typeOfPerson.Field(i).Name
// 		fieldType := typeOfPerson.Field(i).Type

// 		fmt.Printf("%s (%s): %v\n", fieldName, fieldType, field.Interface())

// 	}
// }
