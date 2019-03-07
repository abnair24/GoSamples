package main

import "fmt"

type FullNameFunctionType func(string, string) string

type Name struct{
	FirstName string
	LastName  string
	FullName  FullNameFunctionType
}

func main()  {

	name := Name{
		FirstName:"ABN",
		LastName:"N",
		FullName: func(firstName string, lastName string) string {
			return  firstName +" "+ lastName
		},
	}

	fmt.Println(name.FullName(name.FirstName,name.LastName))


}
