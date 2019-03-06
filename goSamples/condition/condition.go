package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	if 8%4==0 {
		fmt.Println("divisible")
	}

	if num := 9; num<0 {
		fmt.Println(num, "negative")
	} else if num <10 {
		fmt.Println(num, "1 digit")
	} else {
		fmt.Println(num,"multiple	")
	}
}
