package main

import (
	"fmt"
)

func vals() (int,int) {
	return 3, 7
}

func sum(a int, b int) int {
	return a+b
}

func varargsfunc(nums ...int) int {
	sum := 0
	for _,j := range nums {
		sum += j
	}
	return sum
}

func main() {

	a,b := vals()

	fmt.Println(a ," ",b)

	result := sum(3,7)

	fmt.Println(result)

	fmt.Println(varargsfunc(1,2))
	fmt.Println(varargsfunc(10,20,30))

	nums := []int{1,2,3,4,5,6,7,8}

	fmt.Println(varargsfunc(nums...)	)
}

