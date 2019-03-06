package main

import "fmt"

func main() {
	s := make([] string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("s",s)

	s = append(s,"d")
	s = append(s, "f")

	fmt.Println(s)

	l := s[2:4]

	fmt.Println(l)
}
