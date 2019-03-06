package main

import (
	"fmt"

	"sample/user/stringutil"
)

var (
	name string = "Aswathy"
	age int = 32
)

func main() {
	fmt.Println(stringutil.Reverse("Hello world !!!!"))
	name := "Aswathy"
	age := 33

	fmt.Printf("%s (%d) ",name,age)
	fmt.Println(add(43,42))
}

func add(x int, y int) int {
	return x+y
}
