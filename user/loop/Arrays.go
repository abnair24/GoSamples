package main

import "fmt"

func main() {

	var array[5] int
	fmt.Println(array)

	array[4] = 100

	for i:=0 ;i <len(array); i++ {
		fmt.Println(array[i])
	}

	a2 := [6] int {1,2,3,4,5,6}
	fmt.Println(a2)
}
