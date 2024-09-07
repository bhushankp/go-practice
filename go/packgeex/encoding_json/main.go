package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	person1 := Person{Name: "tushar", Email: "tushar@gmail.com", Age: 33}

	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error encoding to json : ", err)
		return
	}
	fmt.Println("encode json data : ", string(jsonData))

	var persn Person

	err = json.Unmarshal(jsonData, &persn) // unmarshaling the same data back into a new variable
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("decoded person %+v :", persn)

}
