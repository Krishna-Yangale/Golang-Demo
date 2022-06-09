package main

import "fmt"

func main() {

	rectangle := Rectangle{
		Length: 20,
		Width:  10,
	}
	fmt.Println("The area of the rectangle is :", rectangle.Area())

	circle := Circle{
		Radius: 10,
	}
	fmt.Println("The area of the circle is:", circle.Area())
}

type Rectangle struct {
	Length int
	Width  int
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
