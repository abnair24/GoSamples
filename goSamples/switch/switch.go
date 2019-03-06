package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {

	case 1: fmt.Println("one")
	case 2: fmt.Println("two")
	case 3: fmt.Println("three")
	}

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday :
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	t := time.Now()

	switch {

	case t.Hour() <12:
		fmt.Println("Morning")

	case t.Hour()> 12 :
		fmt.Println("Evening")
	}


	who := func( i interface{}) {
		switch i.(type) {

		case bool:
			fmt.Println("Boolean")

		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("Stirng")

		default:
			fmt.Println("type unknown")
		}
	}

	who(1)
	who(true)
	who("Aswathy")
}
