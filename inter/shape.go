package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("Площадь фигуры равна = %.2f\n", s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 6}
	circle := Circle{Radius: 78}

	PrintArea(rect)
	PrintArea(circle)
}
