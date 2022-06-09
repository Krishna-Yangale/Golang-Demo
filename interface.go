package main

import "fmt"

func main() {
	Rectangle := Rectangle{10, 30}
	circle := Circle{10}

	details(Rectangle)
	details(circle)

}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func details(s Shape) {
	fmt.Println("Area", s.Area())
}
