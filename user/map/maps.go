package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{"foo":1,"bar":2}

	m1["key1"]= 7
	m1["key2"] = 8

	fmt.Println(m1)

	delete(m1,"key2")

	t,prs := m1["key2"]

	fmt.Println("prs:",prs)
	fmt.Println("t",t)

	fmt.Println(m2)

	rangemap := map[string]string{"a":"apple","b":"banana","c":"cat"}

	for k,v := range rangemap {
		fmt.Println(k," ======= ", v)

	}
}
