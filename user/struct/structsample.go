package main

import (
	"fmt"
)

type Person struct {

	firstname string
	lastname  string
	age 	  int
}

type Vehicle struct {
	Type, Name,Color string
	Weight int
}

func main() {
	var person Person

	fmt.Println("Person initial value: ",person)


	person1 := Person{"Aswathy","Nair",30}

	person = person1

	fmt.Println("person1 :", person1)

	fmt.Println("person :",person)


	car := Vehicle{Type: "Car",
		Name: "Rolls Royce",
		Color: "Red",
		Weight: 1000,
	}

	fmt.Println(car.Name)
	fmt.Println(car.Color)
	fmt.Println(car.Weight)

	car.Color = "Black"

	fmt.Println(car.Color)

	/*
	Structs are value types not reference types
	 */


	p1 := Person{firstname: "ABN",
	lastname:"N",
		age: 1}

	p2 := p1

	fmt.Println("p2 before: ",p2)
	fmt.Println("p1 before: ",p1)

	p1.age = 200

	fmt.Println("p2 after: ",p2)
	fmt.Println("p1 after: ",p1)
}