package geometry

import "math"

// Rectangle struct has a height and a width
type Rectangle struct {
	Height float64
	Width  float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle struct has a radius
type Circle struct {
	Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Shape interface returns Area
type Shape interface {
	Area() float64
}

// Perimeter takes a height and width and returns (h+w)*2
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

// Area takes a height and width and returns h*w
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
