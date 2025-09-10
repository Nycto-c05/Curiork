package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement Shape interface methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Another struct
type Circle struct {
	Radius float64
}

// Implement Shape interface methods for Circle
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func main() {
	// You can hold both Circle and Rectangle in a Shape interface
	var s Shape

	s = Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: Area=%.2f, Perimeter=%.2f\n", s.Area(), s.Perimeter())

	s = Circle{Radius: 7}
	fmt.Printf("Circle: Area=%.2f, Perimeter=%.2f\n", s.Area(), s.Perimeter())
}
