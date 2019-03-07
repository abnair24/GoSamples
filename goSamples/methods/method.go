package main

import (
	"fmt"
	"math"
)

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Update(l float64, w float64) {
	r.Width = w
	r.Length = l
}

func (r *Rectangle) PointerUpdate(l float64 , w float64) {
	r.Width = w
	r.Length = l
}

type Rectangle struct {
	Width float64
	Length float64
}

type Circle struct {
	Radius float64
}

func (d *Details) changePhone(number string) {
	d.phone = number
}

type Contact struct {
	phone string
	address  string
}

type Details struct {
	Name string
	Contact
}

func main() {
	rect := Rectangle{10.0,20.0}
	circle := Circle{5.0}

	fmt.Println(rect.Area())
	fmt.Println(circle.Area())

	rect.Update(5,6)
	fmt.Println("area after :",rect.Area())

	rect.PointerUpdate(5,6)
	fmt.Println("Area after pointer update :",rect.Area())

	details := Details{
		Name: "ABN",
		Contact: Contact{"12314232","sdfsgssdgsgsdsg"},
	}

	fmt.Println("Details :",details)

	details.changePhone("34344343434")
	fmt.Println("New details: ",details)

}