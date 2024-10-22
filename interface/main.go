package main

import "fmt"

// Interface
type Shape interface {
	area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return c.radius * c.radius
}

func main() {
	square := Square{3}
	circle := Circle{10}

	PrintShapeArea(square)
	PrintShapeArea(circle)

}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.area())
}
