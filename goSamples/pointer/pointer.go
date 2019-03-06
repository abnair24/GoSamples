package main

import (
	"fmt"
)



func byVal(val int) {
	val = 0
}

func byRef(ref *int) {
	*ref = 0
}

type Student struct {
	Name string
	Id int
}



func main() {

	i := 10

	fmt.Println("before",i)
	byVal(i)

	fmt.Println("after ",i)

	byRef(&i)

	fmt.Println("after ",i)

	fmt.Println("addr ",&i)


	student := Student{
					Name: "Aswathy",
					Id: 100,
					}

	pointerStudent := &student

	fmt.Println(pointerStudent)

	fmt.Println((*pointerStudent).Id)
	fmt.Println(pointerStudent.Id)
	fmt.Println(pointerStudent.Name)

	pointerStudent.Id = 120

	fmt.Println("student :",student)
	fmt.Println("pointerStudent :",pointerStudent)


	student2 := new(Student)

	fmt.Println("Student 2 :",student2)




	 }
