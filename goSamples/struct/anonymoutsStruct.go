package main

import (
	"fmt"
)

type Salary struct {
	basic int
	insurance int
	allowance int

}
/*
anonymously refering to Salary

 */
type Employee struct {
	firstName string
	lastname string
	id int
	Salary
}

func main() {
	emp := Employee{ firstName: "abn",
	lastname:"n",
	id:1,
//	salary: Salary{}
	Salary : Salary{1000,100,50},
	}

	/*
	For anonymous struct fields can be accessed directly from parent struct using "parent.fieldname" (promoted)
	other using ref, its done : "parent.inneralias.fieldname"
	 */
	emp.basic = 1200

	fmt.Println("basic :",emp.basic)
	fmt.Println("emp :",emp)
}
