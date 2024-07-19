package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	A float64
	B float64
}

func (c Circle) Area() float64 {
	return 3.14 * (math.Pow(2, c.Radius))
}

func (r Rectangle) Area() float64 {
	return r.A * r.B
}

func main () {
	shapeCircle := Circle {
		Radius: 1,
	}

	shapeRectangle := Rectangle {
		A: 1,
		B: 2,
	}
	
	fmt.Println(shapeCircle.Area())
	fmt.Println(shapeRectangle.Area())
}