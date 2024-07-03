package main

import "fmt"

type Shape interface {
	Area() int64
}

type Square struct {
	sideLength int64
}

func (s Square) Area() int64 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius int64
}

func (c Circle) Area() int64 {
	return c.radius * c.radius * 3
}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

func main() {
	square := Square{34}
	circle := Circle{10}

	PrintShapeArea(square)
	PrintShapeArea(circle)
}
