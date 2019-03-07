package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a+b
}

func subtract(a int, b int) int {
	return a-b
}

func calc (a int , b int , fun func(int , int ) int) int {
	result := fun(a,b)
	return result
}
func main()  {

	sum := calc(3,7, add)
	sub := calc(7,3,subtract)

	fmt.Println("sum :",sum)
	fmt.Println("sub :",sub)

}

